package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"rest_service/myservice"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hooklift/gowsdl/soap"
	viteGlue "github.com/torenware/vite-go"
)

type Response struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	OrigError  error  `json:"-"`
}

func (err *Response) Error() string {
	if err.OrigError != nil {
		return err.OrigError.Error()
	}

	return err.Message
}

type OptionalHandler func(http.ResponseWriter, *http.Request) error

func sendResponse(rw http.ResponseWriter, res *Response) {
	fmt.Printf("error: %s\n", res.Error())
	if res.StatusCode == 0 {
		res.StatusCode = 200
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(res.StatusCode)
	json.NewEncoder(rw).Encode(res)
}

func wrapHandler(handler OptionalHandler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := handler(rw, r); err != nil {
			if resp, ok := err.(*Response); ok {
				sendResponse(rw, resp)
			} else {
				sendResponse(rw, &Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Something went wrong.",
					OrigError:  err,
				})
			}
		}
	})
}

type postIdKey struct{}

func getPostId(r *http.Request) int32 {
	return r.Context().Value(postIdKey{}).(int32)
}

func extractPostId(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rawPostId := chi.URLParam(r, "postId")
		postId, err := strconv.Atoi(rawPostId)
		if err != nil {
			sendResponse(rw, &Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid post ID.",
				OrigError:  err,
			})
			return
		}

		handler.ServeHTTP(rw, r.WithContext(context.WithValue(r.Context(), postIdKey{}, int32(postId))))
	})
}

func getPassword(r *http.Request) *string {
	if r.URL.Query().Has("password") {
		gotPassword := r.URL.Query().Get("password")
		return &gotPassword
	}
	return nil
}

func extractRecentPostParams(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		rawLimit := "5"
		if gotRawLimit, hasLimit := query["limit"]; hasLimit && len(gotRawLimit) != 0 {
			rawLimit = gotRawLimit[0]
		}

		rawPage := "1"
		if gotRawPage, hasPage := query["page"]; hasPage && len(gotRawPage) != 0 {
			rawPage = gotRawPage[0]
		}

		limit, err := strconv.Atoi(rawLimit)
		if err != nil || limit < 1 {
			sendResponse(rw, &Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid limit.",
				OrigError:  err,
			})
			return
		}

		ctx := context.WithValue(r.Context(), "limit", limit)
		page, err := strconv.Atoi(rawPage)
		if err != nil || page < 1 {
			sendResponse(rw, &Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid page.",
				OrigError:  err,
			})
			return
		}

		ctx = context.WithValue(ctx, "page", page)
		handler.ServeHTTP(rw, r.WithContext(ctx))
	})
}

type APIService struct {
	Client myservice.IService1
}

func (api *APIService) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(rw, r)
		})
	})
	r.With(extractRecentPostParams).Get("/posts", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
		limit := r.Context().Value("limit").(int)
		page := r.Context().Value("page").(int)
		result, err := api.Client.RecentPosts(&myservice.RecentPosts{
			HowMany: int32(limit),
			Page:    int32(page),
		})
		if err != nil {
			return err
		}
		return json.NewEncoder(rw).Encode(result.RecentPostsResult)
	}))

	r.Post("/posts", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
		post := new(myservice.Post)
		if err := json.NewDecoder(r.Body).Decode(post); err != nil {
			return &Response{
				StatusCode: http.StatusBadRequest,
				Message:    http.StatusText(http.StatusBadRequest),
				OrigError:  err,
			}
		}
		result, err := api.Client.CreatePost(&myservice.CreatePost{
			Post: post,
		})
		if err != nil {
			return err
		} else if !result.CreatePostResult {
			return &Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Post was not created successfully.",
			}
		}

		return &Response{
			StatusCode: http.StatusOK,
			Message:    "Post was created successfully.",
		}
	}))

	r.Route("/posts/{postId}", func(r chi.Router) {
		r.Use(extractPostId)

		r.Get("/is_locked", wrapHandler(func(w http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			result, err := api.Client.IsPostLocked(&myservice.IsPostLocked{
				PostId: postId,
			})
			if err != nil {
				return err
			}
			return json.NewEncoder(w).Encode(result.IsPostLockedResult)
		}))

		r.Get("/", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			password := getPassword(r)
			result, err := api.Client.GetPost(&myservice.GetPost{
				Id:       postId,
				Password: password,
			})
			if err != nil {
				if soapError, ok := err.(*soap.HTTPError); ok {
					statusCode := soapError.StatusCode
					message := "Something went wrong."
					if strings.Contains(string(soapError.ResponseBody), "Post not found!") {
						statusCode = http.StatusNotFound
						message = "Post not found!"
					} else if strings.Contains(string(soapError.ResponseBody), "Password mismatch!") {
						statusCode = http.StatusBadRequest
						message = "Password mismatch!"
					}
					return &Response{
						StatusCode: statusCode,
						Message:    message,
					}
				}
				return err
			}

			return json.NewEncoder(rw).Encode(result.GetPostResult)
		}))

		r.Patch("/", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			password := getPassword(r)
			post := new(myservice.Post)
			if err := json.NewDecoder(r.Body).Decode(post); err != nil {
				return &Response{
					StatusCode: http.StatusBadRequest,
					Message:    http.StatusText(http.StatusBadRequest),
					OrigError:  err,
				}
			}

			if password == nil && post.Password != nil {
				password = post.Password
			}

			post.Id = getPostId(r)
			result, err := api.Client.UpdatePost(&myservice.UpdatePost{
				Post:     post,
				Password: password,
			})
			if err != nil {
				return err
			} else if !result.UpdatePostResult {
				return &Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Post was not updated successfully.",
				}
			}
			return &Response{
				Message: "Post updated successfully.",
			}
		}))

		r.Delete("/", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			password := getPassword(r)
			postId := getPostId(r)
			result, err := api.Client.DeletePost(&myservice.DeletePost{
				Id:       postId,
				Password: password,
			})
			if err != nil {
				return err
			} else if !result.DeletePostResult {
				return &Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Post was not deleted successfully.",
				}
			}
			return &Response{
				Message: "Post was deleted successfully.",
			}
		}))

		r.Get("/likes", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			result, err := api.Client.GetPostLikes(&myservice.GetPostLikes{
				PostId: postId,
			})
			if err != nil {
				return err
			}
			return json.NewEncoder(rw).Encode(result.GetPostLikesResult)
		}))

		r.Get("/is_liked/{sessionId}", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			sessionId := chi.URLParam(r, "sessionId")
			result, err := api.Client.IsPostLiked(&myservice.IsPostLiked{
				PostId:    postId,
				SessionId: &sessionId,
			})
			if err != nil {
				return err
			}
			return json.NewEncoder(rw).Encode(result.IsPostLikedResult)
		}))

		r.Post("/likes/{sessionId}", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			sessionId := chi.URLParam(r, "sessionId")
			result, err := api.Client.LikePost(&myservice.LikePost{
				SessionId: &sessionId,
				PostId:    postId,
			})
			if err != nil {
				return err
			}
			return json.NewEncoder(rw).Encode(result.LikePostResult)
		}))

		r.Delete("/likes/{sessionId}", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			sessionId := chi.URLParam(r, "sessionId")
			result, err := api.Client.UnlikePost(&myservice.UnlikePost{
				SessionId: &sessionId,
				PostId:    postId,
			})
			if err != nil {
				return err
			}
			return json.NewEncoder(rw).Encode(result.UnlikePostResult)
		}))

		r.Get("/views", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			result, err := api.Client.GetPostViews(&myservice.GetPostViews{
				PostId: postId,
			})
			if err != nil {
				return err
			}
			return json.NewEncoder(rw).Encode(result.GetPostViewsResult)
		}))

		r.Post("/views/{sessionId}", wrapHandler(func(rw http.ResponseWriter, r *http.Request) error {
			postId := getPostId(r)
			sessionId := chi.URLParam(r, "sessionId")
			result, err := api.Client.RecordView(&myservice.RecordView{
				SessionId: &sessionId,
				PostId:    postId,
			})
			if err != nil {
				return err
			}
			return json.NewEncoder(rw).Encode(result.RecordViewResult)
		}))
	})

	return r
}

func getEnv(keyName string, defaultValue string) string {
	value, hasKey := os.LookupEnv(keyName)
	if hasKey && len(value) != 0 {
		return value
	}
	return defaultValue
}

//go:embed template.html.tmpl
var templateFile string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error: .env not found")
	}

	wcfServiceUrl := getEnv("SERVICE_URL", "http://localhost:5000/Service1")
	env := getEnv("ENV", "development")
	port := getEnv("PORT", "4003")
	isProd := strings.ToLower(env) == "production"
	frontendFolder := getEnv("FRONTEND_FOLDER", "")
	if len(frontendFolder) == 0 {
		gotFrontendFolder, err := filepath.Abs(filepath.Join("..", "frontend"))
		if err != nil {
			log.Fatalln(err)
		}
		frontendFolder = gotFrontendFolder
	}

	distFolder := frontendFolder
	if isProd {
		distFolder = filepath.Join(frontendFolder, "dist")
	}

	viteConfig := &viteGlue.ViteConfig{
		Environment: env,
		AssetsPath:  ".",
		EntryPoint:  "src/main.ts",
		Platform:    "vue",
		FS:          os.DirFS(distFolder),
	}

	glue, err := viteGlue.NewVueGlue(viteConfig)
	if err != nil {
		log.Fatalln(err)
	}

	ts, err := template.New("index").Parse(templateFile)
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	client := myservice.NewIService1(soap.NewClient(fmt.Sprintf("%s/basichttp", wcfServiceUrl)))
	apiService := APIService{client}
	fsHandler, err := glue.FileServer()
	if err != nil {
		log.Fatalln("could not set up static file server", err)
	}

	assetsEndpointName := "src"
	if isProd {
		assetsEndpointName = "assets"
	}

	r.Handle(fmt.Sprintf("/%s/*", assetsEndpointName), fsHandler)
	r.Mount("/api", apiService.Routes())
	r.Get("/", wrapHandler(func(w http.ResponseWriter, r *http.Request) error {
		return ts.Execute(w, struct {
			Vite      *viteGlue.VueGlue
			PageTitle string
		}{
			Vite:      glue,
			PageTitle: "eFreedomWall",
		})
	}))

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
