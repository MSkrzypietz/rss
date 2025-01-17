<script setup lang="ts">
import PostsAPI, { type Post } from '@/api/posts.ts';
import { OnyxHeadline, OnyxButton } from 'sit-onyx';
import { useFeedStore } from '@/stores/feed.ts';

const feedStore = useFeedStore();

const props = defineProps<{
  post: Post;
}>();

let publishDate: string | null = null;
if (props.post.published_at !== null) {
  const addLeadingZero = (n: number) => n.toString().padStart(2, '0');
  const d = new Date(props.post.published_at);
  const date = `${addLeadingZero(d.getDate())}.${addLeadingZero(d.getMonth() + 1)}.${d.getFullYear()}`;
  const time = `${addLeadingZero(d.getHours())}:${addLeadingZero(d.getMinutes())}:${addLeadingZero(d.getSeconds())}`;
  publishDate = `${date} ${time}`;
}
</script>

<template>
  <a class="feed-item" :href="post.url" target="_blank">
    <div class="feed-item__header">
      <OnyxHeadline is="h2">{{ post.title }}</OnyxHeadline>
      <OnyxButton label="Read" density="compact" @click.stop.prevent="feedStore.markPostAsRead(post.id)"></OnyxButton>
    </div>
    <div class="feed-item__content">
      <span class="onyx-text">{{ post.description }}</span>
    </div>
    <p class="onyx-text">{{ publishDate }} - {{ post.feed_name }}</p>
  </a>
</template>

<style scoped lang="scss">
.feed-item {
  background-color: var(--onyx-color-base-background-blank);
  border-radius: var(--onyx-radius-sm);
  text-decoration: unset;
  color: inherit;
  padding: 1rem;

  &__header {
    display: flex;
    align-items: center;
    gap: var(--onyx-spacing-2xs);

    :deep(.onyx-button) {
      margin-left: auto;
    }
  }

  &__content {
    margin-top: var(--onyx-spacing-2xs);
    margin-bottom: var(--onyx-spacing-4xs);
  }
}
</style>
