<script setup lang="ts">
import { onMounted } from 'vue';
import FeedItem from '@/components/FeedItem.vue';
import { useFeedStore } from '@/stores/feed.ts';

const feedStore = useFeedStore();

onMounted(async () => {
  await feedStore.fetchUnreadPosts();
});
</script>

<template>
  <div class="onyx-grid">
    <FeedItem v-for="post in feedStore.posts" :post="post" class="onyx-grid-span-16">{{ post }}</FeedItem>
  </div>
</template>

<style scoped lang="scss"></style>
