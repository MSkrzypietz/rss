<script setup lang="ts">
import { onMounted } from 'vue';
import FeedItem from '@/components/FeedItem.vue';
import { useFeedStore } from '@/stores/feed.ts';
import { OnyxButton } from 'sit-onyx';

const feedStore = useFeedStore();

onMounted(async () => {
  await feedStore.fetchUnreadPosts();
});
</script>

<template>
  <div class="onyx-grid feed-item-list">
    <FeedItem v-for="post in feedStore.posts" :post="post" class="onyx-grid-span-16">{{ post }}</FeedItem>
    <OnyxButton class="onyx-grid-span-16" label="Refresh" @click="feedStore.fetchUnreadPosts()" />
  </div>
</template>

<style scoped lang="scss">
.feed-item-list {
  .onyx-button {
    margin: auto;
  }
}
</style>
