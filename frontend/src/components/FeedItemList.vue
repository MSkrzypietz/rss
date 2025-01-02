<script setup lang="ts">
import { onMounted, ref } from 'vue';
import PostsAPI, { type Post } from '@/api/posts.ts';
import FeedItem from '@/components/FeedItem.vue';

let posts = ref(null as Post[] | null);

onMounted(async () => {
  posts.value = await PostsAPI.getUnreadPosts();
});
</script>

<template>
  FeedItemList
  <FeedItem v-for="post in posts" :post="post">{{ post }}</FeedItem>
</template>

<style scoped lang="scss"></style>
