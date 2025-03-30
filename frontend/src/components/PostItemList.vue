<script setup lang="ts">
import { onMounted, watch } from 'vue';
import PostItem from '@/components/PostItem.vue';
import { usePostStore } from '@/stores/post.ts';
import { OnyxButton } from 'sit-onyx';
import SearchBar from '@/components/SearchBar.vue';
import { debounce } from '@/utils/debounce.ts';
import { useFeedStore } from '@/stores/feed.ts';

const postStore = usePostStore();
const feedStore = useFeedStore();

onMounted(async () => {
  await Promise.all([postStore.fetchUnreadPosts(), feedStore.fetchFeeds()]);
});

watch(
  postStore.postFilter,
  debounce(() => postStore.fetchUnreadPosts(), 500),
);
</script>

<template>
  <SearchBar />
  <div class="onyx-grid post-item-list">
    <PostItem v-for="post in postStore.posts" :post="post" class="onyx-grid-span-16">{{ post }}</PostItem>
    <OnyxButton class="onyx-grid-span-16" label="Refresh" @click="postStore.fetchUnreadPosts()" />
  </div>
</template>

<style scoped lang="scss">
.post-item-list {
  margin-top: var(--onyx-spacing-lg);

  .onyx-button {
    margin: auto;
  }
}
</style>
