<script setup lang="ts">
import { onMounted, watch } from 'vue';
import PostItem from '@/components/PostItem.vue';
import { usePostStore } from '@/stores/post.ts';
import { OnyxButton, OnyxSkeleton } from 'sit-onyx';
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
  <div>
    <SearchBar />
    <div class="onyx-grid post-item-list">
      <OnyxSkeleton v-if="postStore.isFetchingPosts" v-for="_ in 10" class="onyx-grid-span-16" />
      <PostItem v-else v-for="post in postStore.posts" :post="post" class="onyx-grid-span-16">{{ post }}</PostItem>
      <OnyxButton class="onyx-grid-span-16" label="Refresh" @click="postStore.fetchUnreadPosts()" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.post-item-list {
  margin-top: var(--onyx-spacing-lg);

  .onyx-skeleton {
    height: 150px;
  }

  .onyx-button {
    margin: auto;
  }
}
</style>
