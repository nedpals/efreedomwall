<template>
    <div class="hero -mt-24">
      <div class="content-container text-center">
        <h1>Post whatever is on your mind. Anonymously.</h1>
        <p>Share your feelings. Confess to someone. Post anything without an account with eFreedomWall.</p>
      </div>
    </div>

    <div class="content-container main-content">
      <post-form @success="() => refetch()" />

      <section id="posts" class="mt-8">
          <h2 class="text-4xl font-bold mb-4">Recent Posts</h2>
          <span v-if="isLoading">Loading...</span>
          <div v-if="postList" class="flex flex-wrap">
            <template  :key="'posts_key_' + i" v-for="(posts, i) in postList.pages">
              <template :key="'post_key_' + j" v-for="(post, j) in posts.Results">
                <div class="w-1/3 p-1">
                  <router-link
                    class="box p-6 is-hoverable hover:shadow-lg relative h-full flex flex-col"
                    :to="{ name: 'post-page', params: { post_id: post.Id } }">
                    <div class="min-h-[7rem] flex flex-col">
                      <p class="text-xl italic font-semibold">
                        {{ post.IsLocked ? 'Post locked.' : post.Content }}
                      </p>
                    </div>
                    
                    <div class="text-gray-500 flex justify-between">
                      <span class="font-bold truncate max-w-[6.3rem]" v-if="!post.IsLocked">~ {{ post.Poster || 'anon' }}</span>
                      <span class="ml-auto">{{ humanizeTime(post.CreatedAt) }}</span>
                    </div>
                  </router-link>
                </div>
              </template>
            </template>
          </div>
          <button 
            class="mt-4 button is-secondary w-full"
            v-if="!isFetchingNextPage && hasNextPage" @click="() => fetchNextPage()">
            Next
          </button>
      </section>
    </div>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';
import { Post, Posts } from '../types';
import { client } from '../client';
import { useInfiniteQuery } from 'vue-query';
import { humanizeTime } from '../utils';
import PostForm from '../components/PostForm.vue';

function useRecentPostsQuery() {
  return useInfiniteQuery<Posts>(['recentPosts'], async ({ pageParam = 1 }) => {
    const resp = await client(`/posts?page=${pageParam}&limit=6`);
    return resp.json()
  }, {
    refetchOnWindowFocus: false,
    getNextPageParam: (lastPage, pages) => lastPage.NextPage == 0 ? null : lastPage.NextPage,
    getPreviousPageParam: (lastPage, pages) => lastPage.PrevPage == 0 ? null : lastPage.PrevPage,
    select: ({ pages, pageParams }) => {
      return {
        pages: pages.map(fetchedPosts => ({
          ...fetchedPosts,
          Results: fetchedPosts.Results.map<Post>((p: any) => ({
              Id: p.Id,
              Content: p.Content,
              Poster: p.Poster,
              IsLocked: p.IsLocked,
              CreatedAt: new Date(p.CreatedAt)
          }))
        })),
        pageParams
      };
    },
    keepPreviousData: true,
  });
}

const { data: postList, refetch, isLoading, hasNextPage, isFetchingNextPage, fetchNextPage } = useRecentPostsQuery();
</script>

<style lang="postcss">
.home_page .navbar {
  background: none !important;
  @apply py-8;
}

.home_page .hero {
  @apply px-[0.8rem] pt-35 pb-40 text-white bg-blue-800 bg-gradient-to-tl to-blue-800 from-pink-300;
}

.home_page .hero > .content-container h1 {
  @apply text-6xl font-bold mb-8;
}

.home_page .hero > .content-container p {
  @apply text-3xl;
}

.home_page .main-content {
  @apply -mt-24;
}

.home_page textarea {
  @apply outline-none p-0 text-3xl min-h-40;
}

.home_page #posts .post-list {
  @apply -mx-1;
}
</style>