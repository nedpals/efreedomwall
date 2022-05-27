<template>
  <form class="box with-padding post-form" @submit.prevent="submitForm($event as SubmitEvent)">
    <div class="field">
        <textarea 
          placeholder="What's on your mind?" 
          name="post_content" 
          id="post_content"
          :value="post?.Content ?? ''"></textarea>
    </div>
    <div class="field-group">
      <div class="field flex-1">
          <label for="post_poster">Poster</label>
          <input type="text" name="post_poster" id="post_poster" :value="post?.Poster ?? 'Anonymous'" />
      </div>
      <div class="field flex-1">
          <label for="post_password">Password <span class="text-gray-500">(Optional)</span> </label>
          <input type="password" name="post_password" id="post_password" />
      </div>
    </div>
    <button class="button is-primary self-end mt-2 px-8" type="submit">{{ post ? 'Update' : 'Post' }}</button>
  </form>
</template>

<script lang="ts" setup>
import { defineEmits, defineProps, PropType } from 'vue';
import { client } from '../client';
import { Post } from '../types';

const emit = defineEmits(['success']);
const props = defineProps({
  post: {
    type: Object as PropType<Post>,
  }
});

async function submitForm(e: SubmitEvent) {
    if (!(e.target instanceof HTMLFormElement))
        return;

    const endpoint = props.post ? `/posts/${props.post.Id}` : '/posts';
    const endpointMethod = props.post ? 'PATCH' : 'POST';
    const form = e.target;
    const data = new FormData(form);
    const content = data.get('post_content');
    const poster = data.get('post_poster');
    const password = data.get('post_password');
    const resp = await client(endpoint, {
        method: endpointMethod,
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            Content: content,
            Poster: poster,
            Password: props.post || !password || password.toString().length == 0 ? null : password.toString()
        })
    });

    if (resp.ok) {
      emit('success');
      form.reset();
    }
}
</script>

<style lang="postcss" scoped>
.post-form {
  @apply flex flex-col shadow-lg border-0 items-start;
}

.post-form .field {
  @apply my-1 flex flex-col w-full;
}

.post-form .field-group {
  @apply flex w-full;
}

.post-form .field-group > .field {
  @apply p-2 w-full;
}

.post-form .field-group > .field:first-child {
  @apply pl-0;
}

.post-form .field-group > .field:last-child {
  @apply pr-0;
}

.post-form label {
  @apply mb-2;
}
</style>