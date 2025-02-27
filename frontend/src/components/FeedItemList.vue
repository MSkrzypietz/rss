<script setup lang="ts">
import { onMounted, watch } from 'vue';
import FeedItem from '@/components/FeedItem.vue';
import { useFeedStore } from '@/stores/feed.ts';
import { OnyxButton } from 'sit-onyx';
import SearchBar from '@/components/SearchBar.vue';
import { debounce } from '@/utils/debounce.ts';

const feedStore = useFeedStore();

onMounted(async () => {
  await Promise.all([feedStore.fetchUnreadPosts(), feedStore.fetchFeeds()]);
});

watch(
  feedStore.postFilter,
  debounce(() => feedStore.fetchUnreadPosts(), 500),
);
</script>

<template>
  <SearchBar />
  <div class="onyx-grid feed-item-list">
    <FeedItem v-for="post in feedStore.posts" :post="post" class="onyx-grid-span-16">{{ post }}</FeedItem>
    <OnyxButton class="onyx-grid-span-16" label="Refresh" @click="feedStore.fetchUnreadPosts()" />
  </div>
</template>

<style scoped lang="scss">
.feed-item-list {
  margin-top: var(--onyx-spacing-lg);

  .onyx-button {
    margin: auto;
  }
}
</style>
