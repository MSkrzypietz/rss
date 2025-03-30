<script setup lang="ts">
import { computed } from 'vue';
import { OnyxDataGrid, type ColumnConfig, type ColumnGroupConfig, OnyxHeadline } from 'sit-onyx';
import { useFeedStore } from '@/stores/feed.ts';
import type { Feed } from '@/api/feeds.ts';

const feedStore = useFeedStore();

const data = computed(() => {
  return feedStore.feeds ?? [];
});

const columns: ColumnConfig<Feed, ColumnGroupConfig, never>[] = [
  { key: 'name', label: 'Name' },
  { key: 'url', label: 'URL' },
];
</script>

<template>
  <div>
    <OnyxHeadline is="h2">All Feeds</OnyxHeadline>
    <OnyxDataGrid :columns :data />
  </div>
</template>

<style scoped lang="scss">
.onyx-data-grid {
  margin-top: var(--onyx-spacing-lg);
}
</style>
