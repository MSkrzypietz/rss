<script setup lang="ts">
import FeedItemList from '@/components/FeedItemList.vue';
import { OnyxPageLayout } from 'sit-onyx';
</script>

<template>
  <OnyxPageLayout>
    <div class="onyx-grid-container">
      <h1>Feed</h1>
      <FeedItemList />
    </div>
  </OnyxPageLayout>
</template>

<style scoped lang="scss"></style>
