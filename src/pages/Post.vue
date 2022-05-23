<template>
    <div class="content-container py-8">
        <article class="box">
            <div v-if="isLockedLoading || isLoading" class="py-12 min-h-xl flex flex-col items-center justify-center">
                <p class="text-3xl">Loading...</p>
            </div>
            
            <form 
                v-else-if="!unlocked && isLocked" 
                class="px-8 py-12 text-center flex flex-col"
                @submit.prevent="(e) => enterPost(e as SubmitEvent)">
                <div class="flex flex-col items-center text-center mb-8">
                    <icon-lock class="text-10xl text-gray-400 mb-4" />
                    <h2 class="text-4xl mb-3 font-semibold">This post is locked.</h2>
                    <p class="text-xl">To access this post, please enter the password.</p>
                </div>

                <input type="password" class="self-center w-2/3 mb-4" name="password" id="password">
                <button class="self-center w-1/2 button is-primary" type="submit">Unlock</button>
            </form>

            <template v-else-if="unlocked && post">
                <section class="flex text-center items-center justify-center bg-gradient-4 text-white rounded-t-lg min-h-lg">
                    <p class="text-4xl font-bold">{{ post.Content }}</p>
                </section>
                <section class="text-gray-500 rounded-b-lg flex px-6 py-4 w-full space-x-4">
                    <p class="font-bold mr-auto" v-if="post?.Poster">
                        <span class="font-normal">by</span> ~{{ post.Poster }}
                    </p>
                    <div class="flex items-center space-x-1">
                        <icon-clock />
                        <span>{{ humanizeTime(post.CreatedAt) }}</span>
                    </div>
                    <div class="flex items-center space-x-1">
                        <icon-view />
                        <span>{{ postViews }} views</span>
                    </div>
                    <button
                        class="hover:text-gray-700 flex items-center space-x-1"
                        @click="isPostLiked ? unlikePost({ sessionId }) : likePost({ sessionId })">
                        <icon-like :class="{ 'text-pink-600': isPostLiked }" />
                        <span>{{ postLikes && postLikes > 0 ? postLikes + ' ' + (postLikes == 1 ? 'Like' : 'Likes') : 'Like' }}</span>
                    </button>
                    <button
                        class="text-gray-700 group flex items-center space-x-1"
                        @click="() => deletePost({ password: storedPassword })">
                        <icon-delete class="group-hover:text-red-700" />
                        <span>Delete</span>
                    </button>
                </section>
            </template>
        </article>
    </div>
</template>

<script lang="ts" setup>
import IconLock from '~icons/uil/lock-alt';
import IconDelete from '~icons/uil/trash-alt';
import IconLike from '~icons/uil/heart';
import IconClock from '~icons/uil/clock';
import IconView from '~icons/uil/eye';

import { useRoute, useRouter } from 'vue-router';
import { client } from '../client';
import { useQuery, useMutation } from 'vue-query';
import { ref } from 'vue';
import { generateSessionId, sessionId } from '../session';
import { humanizeTime } from '../utils';

const router = useRouter();
const route = useRoute();
const unlocked = ref(false);
const storedPassword = ref("");

function usePostIsLikedQuery(id: string, sessionId: string) {
    return useQuery<boolean>(['posts/likes', id, sessionId], async () => {
        const resp = await client(`/posts/${id}/is_liked/${sessionId}`);
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return resp.json();
    }, {
        enabled: unlocked,
        refetchOnWindowFocus: false
    })
}

function usePostLikesQuery(id: string) {
    return useQuery<number>(['posts/likes', id], async () => {
        const resp = await client(`/posts/${id}/likes`);
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return resp.json();
    }, {
        initialData: 0,
        enabled: unlocked
    })
}

function useUnlikePostMutation(id: string) {
    return useMutation(async ({ sessionId }: { sessionId: string }) => {
        const resp = await client(`/posts/${id}/likes/${sessionId}`, {
            method: 'DELETE'
        });
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return resp.json();
    }, {
        onSuccess: (data) => {
            refetchPostLikes();
            refetchIsPostLiked();
        }
    })
}

function useLikePostMutation(id: string) {
    return useMutation(async ({ sessionId }: { sessionId: string }) => {
        const resp = await client(`/posts/${id}/likes/${sessionId}`, {
            method: 'POST'
        });
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return resp.json();
    }, {
        onSuccess: (data) => {
            refetchPostLikes();
            refetchIsPostLiked();
        }
    })
}

function usePostViewsMutation(id: string) {
    return useMutation(async ({ sessionId }: { sessionId: string }) => {
        const resp = await client(`/posts/${id}/views/${sessionId}`, {
            method: 'POST'
        });
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return resp.json();
    }, {
        onSuccess: () => {
            refetchPostViews();
        }
    })
}

function usePostViewsQuery(id: string) {
    return useQuery<number>(['posts/views', id], async () => {
        const resp = await client(`/posts/${id}/views`);
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return resp.json();
    }, {
        initialData: 0,
        enabled: unlocked
    })
}

function useIsLockedQuery(id: string) {
    return useQuery<boolean>(['posts/isLocked', id], async () => {
        const isLockedResp = await client(`/posts/${id}/is_locked`);
        if (!isLockedResp.ok) {
            throw new Error(await isLockedResp.text());
        }
        return isLockedResp.json();
    }, {
        refetchOnWindowFocus: false,
        onSuccess: (isLocked) => {
            if (!isLocked) {
                unlockPost({ password: '' });
            }
        }
    });
}

function useGetPostMutation(id: string) {
    return useMutation(async ({ password = "" } : {password: string}) => {
        const resp = await client(`/posts/${id}?password=${password}`);
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return await resp.json();
    }, {
        onSuccess: () => {
            recordView({ sessionId: sessionId.value });
            unlocked.value = true;
        }
    });
}

function useDeletePostMutation(id: string) {
    return useMutation(async ({ password = "" } : {password: string}) => {
        const resp = await client(`/posts/${id}?password=${password}`, {
            method: 'DELETE'
        });
        if (!resp.ok) {
            throw new Error(await resp.text());
        }
        return await resp.json();
    }, {
        onSuccess: () => {
            // TODO: add message
            router.replace({ name: 'home' });
        }
    });
}

function enterPost(ev: SubmitEvent) {
    if (!(ev.target instanceof HTMLFormElement))
        return;

    const data = new FormData(ev.target);
    storedPassword.value = data.get('password')?.toString() ?? ''
    unlockPost({ password: storedPassword.value });
}

const postId = route.params.post_id as string;
const { data: postViews, refetch: refetchPostViews } = usePostViewsQuery(postId);
const { data: postLikes, refetch: refetchPostLikes } = usePostLikesQuery(postId);
const { data: isPostLiked, refetch: refetchIsPostLiked } = usePostIsLikedQuery(postId, sessionId.value ?? generateSessionId());
const { data: isLocked, isLoading: isLockedLoading } = useIsLockedQuery(postId);
const { mutate: unlockPost, data: post, isLoading } = useGetPostMutation(postId);
const { mutate: recordView } = usePostViewsMutation(postId);
const { mutate: likePost } = useLikePostMutation(postId);
const { mutate: unlikePost } = useUnlikePostMutation(postId);
const { mutate: deletePost } = useDeletePostMutation(postId);
</script>