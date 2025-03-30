<script setup lang="ts">
import { OnyxPageLayout } from 'sit-onyx';
import FeedItemNew from '@/components/FeedItemNew.vue';
import FeedItemList from '@/components/FeedItemList.vue';
</script>

<template>
  <OnyxPageLayout>
    <div class="onyx-grid-container page-container">
      <FeedItemNew class="my-content" />
      <FeedItemList class="my-content feed-item-list" />
    </div>
  </OnyxPageLayout>
</template>

<style scoped lang="scss">
.page-container {
  display: flex;
  flex-direction: column;
  gap: var(--onyx-spacing-2xl);
}
</style>
