// Code generated by gowsdl DO NOT EDIT.

package myservice

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type Post struct {
	Content *string `xml:"http://schemas.datacontract.org/2004/07/EFreedomWallService Content,omitempty" json:"Content,omitempty"`

	CreatedAt DateTime `xml:"http://schemas.datacontract.org/2004/07/EFreedomWallService CreatedAt,omitempty" json:"CreatedAt,omitempty"`

	Id int32 `xml:"http://schemas.datacontract.org/2004/07/EFreedomWallService Id,omitempty" json:"Id,omitempty"`
	
	IsLocked bool `xml:"http://schemas.datacontract.org/2004/07/EFreedomWallService IsLocked,omitempty" json:"IsLocked,omitempty"`

	Password *string `xml:"http://schemas.datacontract.org/2004/07/EFreedomWallService Password,omitempty" json:"Password,omitempty"`

	Poster *string `xml:"http://schemas.datacontract.org/2004/07/EFreedomWallService Poster,omitempty" json:"Poster,omitempty"`
}

type Posts struct {
	CurrentPage int32 `xml:"CurrentPage,omitempty" json:"CurrentPage,omitempty"`

	NextPage *int32 `xml:"NextPage,omitempty" json:"NextPage,omitempty"`

	PrevPage *int32 `xml:"PrevPage,omitempty" json:"PrevPage,omitempty"`

	Results *ArrayOfPost `xml:"Results,omitempty" json:"Results,omitempty"`

	Total int32 `xml:"Total,omitempty" json:"Total,omitempty"`
}

type ArrayOfPost struct {
	Post []*Post `xml:"Post,omitempty" json:"Post,omitempty"`
}

type IsPostLocked struct {
	XMLName xml.Name `xml:"http://tempuri.org/ IsPostLocked"`

	PostId int32 `xml:"postId,omitempty" json:"postId,omitempty"`
}

type IsPostLockedResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ IsPostLockedResponse"`

	IsPostLockedResult bool `xml:"IsPostLockedResult,omitempty" json:"IsPostLockedResult,omitempty"`
}

type CreatePost struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreatePost"`

	Post *Post `xml:"post,omitempty" json:"post,omitempty"`
}

type CreatePostResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreatePostResponse"`

	CreatePostResult bool `xml:"CreatePostResult,omitempty" json:"CreatePostResult,omitempty"`
}

type GetPost struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetPost"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`

	Password *string `xml:"password,omitempty" json:"password,omitempty"`
}

type GetPostResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetPostResponse"`

	GetPostResult *Post `xml:"GetPostResult,omitempty" json:"GetPostResult,omitempty"`
}

type UpdatePost struct {
	XMLName xml.Name `xml:"http://tempuri.org/ UpdatePost"`

	Post *Post `xml:"post,omitempty" json:"post,omitempty"`

	Password *string `xml:"password,omitempty" json:"password,omitempty"`
}

type UpdatePostResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ UpdatePostResponse"`

	UpdatePostResult bool `xml:"UpdatePostResult,omitempty" json:"UpdatePostResult,omitempty"`
}

type DeletePost struct {
	XMLName xml.Name `xml:"http://tempuri.org/ DeletePost"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`

	Password *string `xml:"password,omitempty" json:"password,omitempty"`
}

type DeletePostResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ DeletePostResponse"`

	DeletePostResult bool `xml:"DeletePostResult,omitempty" json:"DeletePostResult,omitempty"`
}

type RecordView struct {
	XMLName xml.Name `xml:"http://tempuri.org/ RecordView"`

	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty"`

	PostId int32 `xml:"postId,omitempty" json:"postId,omitempty"`
}

type RecordViewResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ RecordViewResponse"`

	RecordViewResult int32 `xml:"RecordViewResult,omitempty" json:"RecordViewResult,omitempty"`
}

type LikePost struct {
	XMLName xml.Name `xml:"http://tempuri.org/ LikePost"`

	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty"`

	PostId int32 `xml:"postId,omitempty" json:"postId,omitempty"`
}

type LikePostResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ LikePostResponse"`

	LikePostResult int32 `xml:"LikePostResult,omitempty" json:"LikePostResult,omitempty"`
}

type UnlikePost struct {
	XMLName xml.Name `xml:"http://tempuri.org/ UnlikePost"`

	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty"`

	PostId int32 `xml:"postId,omitempty" json:"postId,omitempty"`
}

type UnlikePostResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ UnlikePostResponse"`

	UnlikePostResult int32 `xml:"UnlikePostResult,omitempty" json:"UnlikePostResult,omitempty"`
}

type IsPostLiked struct {
	XMLName xml.Name `xml:"http://tempuri.org/ IsPostLiked"`

	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty"`

	PostId int32 `xml:"postId,omitempty" json:"postId,omitempty"`
}

type IsPostLikedResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ IsPostLikedResponse"`

	IsPostLikedResult bool `xml:"IsPostLikedResult,omitempty" json:"IsPostLikedResult,omitempty"`
}

type GetPostLikes struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetPostLikes"`

	PostId int32 `xml:"postId,omitempty" json:"postId,omitempty"`
}

type GetPostLikesResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetPostLikesResponse"`

	GetPostLikesResult int32 `xml:"GetPostLikesResult,omitempty" json:"GetPostLikesResult,omitempty"`
}

type GetPostViews struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetPostViews"`

	PostId int32 `xml:"postId,omitempty" json:"postId,omitempty"`
}

type GetPostViewsResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetPostViewsResponse"`

	GetPostViewsResult int32 `xml:"GetPostViewsResult,omitempty" json:"GetPostViewsResult,omitempty"`
}

type RecentPosts struct {
	XMLName xml.Name `xml:"http://tempuri.org/ RecentPosts"`

	HowMany int32 `xml:"howMany,omitempty" json:"howMany,omitempty"`

	Page int32 `xml:"page,omitempty" json:"page,omitempty"`
}

type RecentPostsResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ RecentPostsResponse"`

	RecentPostsResult *Posts `xml:"RecentPostsResult,omitempty" json:"RecentPostsResult,omitempty"`
}

type Char int32

type Guid string

type Base64Binary []byte

type Boolean bool

type DateTime soap.XSDDateTime

func (xdt DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return soap.XSDDateTime(xdt).MarshalXML(e, start)
}

func (xdt *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*soap.XSDDateTime)(xdt).UnmarshalXML(d, start)
}

type Decimal float64

type Double float64

type Float float32

type int int32

type Long int64

type Short int16

type UnsignedByte byte

type UnsignedInt uint32

type UnsignedLong uint64

type UnsignedShort uint16

type IService1 interface {
	IsPostLocked(request *IsPostLocked) (*IsPostLockedResponse, error)

	IsPostLockedContext(ctx context.Context, request *IsPostLocked) (*IsPostLockedResponse, error)

	CreatePost(request *CreatePost) (*CreatePostResponse, error)

	CreatePostContext(ctx context.Context, request *CreatePost) (*CreatePostResponse, error)

	GetPost(request *GetPost) (*GetPostResponse, error)

	GetPostContext(ctx context.Context, request *GetPost) (*GetPostResponse, error)

	UpdatePost(request *UpdatePost) (*UpdatePostResponse, error)

	UpdatePostContext(ctx context.Context, request *UpdatePost) (*UpdatePostResponse, error)

	DeletePost(request *DeletePost) (*DeletePostResponse, error)

	DeletePostContext(ctx context.Context, request *DeletePost) (*DeletePostResponse, error)

	RecordView(request *RecordView) (*RecordViewResponse, error)

	RecordViewContext(ctx context.Context, request *RecordView) (*RecordViewResponse, error)

	LikePost(request *LikePost) (*LikePostResponse, error)

	LikePostContext(ctx context.Context, request *LikePost) (*LikePostResponse, error)

	UnlikePost(request *UnlikePost) (*UnlikePostResponse, error)

	UnlikePostContext(ctx context.Context, request *UnlikePost) (*UnlikePostResponse, error)

	IsPostLiked(request *IsPostLiked) (*IsPostLikedResponse, error)

	IsPostLikedContext(ctx context.Context, request *IsPostLiked) (*IsPostLikedResponse, error)

	GetPostLikes(request *GetPostLikes) (*GetPostLikesResponse, error)

	GetPostLikesContext(ctx context.Context, request *GetPostLikes) (*GetPostLikesResponse, error)

	GetPostViews(request *GetPostViews) (*GetPostViewsResponse, error)

	GetPostViewsContext(ctx context.Context, request *GetPostViews) (*GetPostViewsResponse, error)

	RecentPosts(request *RecentPosts) (*RecentPostsResponse, error)

	RecentPostsContext(ctx context.Context, request *RecentPosts) (*RecentPostsResponse, error)
}

type iService1 struct {
	client *soap.Client
}

func NewIService1(client *soap.Client) IService1 {
	return &iService1{
		client: client,
	}
}

func (service *iService1) IsPostLockedContext(ctx context.Context, request *IsPostLocked) (*IsPostLockedResponse, error) {
	response := new(IsPostLockedResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/IsPostLocked", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) IsPostLocked(request *IsPostLocked) (*IsPostLockedResponse, error) {
	return service.IsPostLockedContext(
		context.Background(),
		request,
	)
}

func (service *iService1) CreatePostContext(ctx context.Context, request *CreatePost) (*CreatePostResponse, error) {
	response := new(CreatePostResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/CreatePost", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) CreatePost(request *CreatePost) (*CreatePostResponse, error) {
	return service.CreatePostContext(
		context.Background(),
		request,
	)
}

func (service *iService1) GetPostContext(ctx context.Context, request *GetPost) (*GetPostResponse, error) {
	response := new(GetPostResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/GetPost", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) GetPost(request *GetPost) (*GetPostResponse, error) {
	return service.GetPostContext(
		context.Background(),
		request,
	)
}

func (service *iService1) UpdatePostContext(ctx context.Context, request *UpdatePost) (*UpdatePostResponse, error) {
	response := new(UpdatePostResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/UpdatePost", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) UpdatePost(request *UpdatePost) (*UpdatePostResponse, error) {
	return service.UpdatePostContext(
		context.Background(),
		request,
	)
}

func (service *iService1) DeletePostContext(ctx context.Context, request *DeletePost) (*DeletePostResponse, error) {
	response := new(DeletePostResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/DeletePost", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) DeletePost(request *DeletePost) (*DeletePostResponse, error) {
	return service.DeletePostContext(
		context.Background(),
		request,
	)
}

func (service *iService1) RecordViewContext(ctx context.Context, request *RecordView) (*RecordViewResponse, error) {
	response := new(RecordViewResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/RecordView", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) RecordView(request *RecordView) (*RecordViewResponse, error) {
	return service.RecordViewContext(
		context.Background(),
		request,
	)
}

func (service *iService1) LikePostContext(ctx context.Context, request *LikePost) (*LikePostResponse, error) {
	response := new(LikePostResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/LikePost", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) LikePost(request *LikePost) (*LikePostResponse, error) {
	return service.LikePostContext(
		context.Background(),
		request,
	)
}

func (service *iService1) UnlikePostContext(ctx context.Context, request *UnlikePost) (*UnlikePostResponse, error) {
	response := new(UnlikePostResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/UnlikePost", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) UnlikePost(request *UnlikePost) (*UnlikePostResponse, error) {
	return service.UnlikePostContext(
		context.Background(),
		request,
	)
}

func (service *iService1) IsPostLikedContext(ctx context.Context, request *IsPostLiked) (*IsPostLikedResponse, error) {
	response := new(IsPostLikedResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/IsPostLiked", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) IsPostLiked(request *IsPostLiked) (*IsPostLikedResponse, error) {
	return service.IsPostLikedContext(
		context.Background(),
		request,
	)
}

func (service *iService1) GetPostLikesContext(ctx context.Context, request *GetPostLikes) (*GetPostLikesResponse, error) {
	response := new(GetPostLikesResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/GetPostLikes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) GetPostLikes(request *GetPostLikes) (*GetPostLikesResponse, error) {
	return service.GetPostLikesContext(
		context.Background(),
		request,
	)
}

func (service *iService1) GetPostViewsContext(ctx context.Context, request *GetPostViews) (*GetPostViewsResponse, error) {
	response := new(GetPostViewsResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/GetPostViews", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) GetPostViews(request *GetPostViews) (*GetPostViewsResponse, error) {
	return service.GetPostViewsContext(
		context.Background(),
		request,
	)
}

func (service *iService1) RecentPostsContext(ctx context.Context, request *RecentPosts) (*RecentPostsResponse, error) {
	response := new(RecentPostsResponse)
	err := service.client.CallContext(ctx, "http://tempuri.org/IService1/RecentPosts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iService1) RecentPosts(request *RecentPosts) (*RecentPostsResponse, error) {
	return service.RecentPostsContext(
		context.Background(),
		request,
	)
}
