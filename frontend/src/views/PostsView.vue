<script setup lang="ts">
import PostItemList from '@/components/PostItemList.vue';
import { OnyxPageLayout } from 'sit-onyx';
</script>

<template>
  <OnyxPageLayout>
    <div class="onyx-grid-container">
      <PostItemList />
    </div>
  </OnyxPageLayout>
</template>

<style scoped lang="scss"></style>
