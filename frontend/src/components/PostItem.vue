<script setup lang="ts">
import { type Post } from '@/api/posts.ts';
import { OnyxHeadline, OnyxButton } from 'sit-onyx';
import { usePostStore } from '@/stores/post.ts';
import { computed } from 'vue';

const postStore = usePostStore();

const props = defineProps<{
  post: Post;
}>();

const publishDate = computed(() => {
  if (props.post.published_at === null) {
    return null;
  }
  const addLeadingZero = (n: number) => n.toString().padStart(2, '0');
  const d = new Date(props.post.published_at);
  const date = `${addLeadingZero(d.getDate())}.${addLeadingZero(d.getMonth() + 1)}.${d.getFullYear()}`;
  const time = `${addLeadingZero(d.getHours())}:${addLeadingZero(d.getMinutes())}:${addLeadingZero(d.getSeconds())}`;
  return `${date} ${time}`;
});

const descriptionLength = 500;
const description = computed(() => {
  const desc = props.post.description.trim();
  if (desc.length <= descriptionLength) {
    return desc;
  }
  return desc.slice(0, descriptionLength) + '...';
});
</script>

<template>
  <a class="post-item" :href="post.url" target="_blank">
    <div class="post-item__header">
      <OnyxHeadline is="h2">{{ post.title }}</OnyxHeadline>
      <OnyxButton label="Read" density="compact" @click.stop.prevent="postStore.markPostAsRead(post.id)"></OnyxButton>
    </div>
    <div class="post-item__content">
      <span class="onyx-text">{{ description }}</span>
    </div>
    <p class="onyx-text">{{ publishDate }} - {{ post.feed_name }}</p>
  </a>
</template>

<style scoped lang="scss">
.post-item {
  background-color: var(--onyx-color-base-background-blank);
  border-radius: var(--onyx-radius-sm);
  text-decoration: unset;
  color: inherit;
  padding: 1rem;

  &:hover {
    box-shadow: 0 0 10px 5px rgba(0, 150, 255, 0.7);
  }

  &__header {
    display: flex;
    align-items: center;
    gap: var(--onyx-spacing-2xs);

    .onyx-button {
      margin-left: auto;
    }
  }

  &__content {
    margin-top: var(--onyx-spacing-2xs);
    margin-bottom: var(--onyx-spacing-4xs);
  }
}
</style>
