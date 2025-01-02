<script setup lang="ts">
import { type Post } from '@/api/posts.ts';

const props = defineProps<{
  post: Post;
}>();
</script>

<template>
  {{ post.title }}
  {{ post.description }}
</template>

<style scoped lang="scss"></style>
