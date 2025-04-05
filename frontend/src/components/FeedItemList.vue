<script setup lang="ts">
import { computed } from 'vue';
import { OnyxDataGrid, type ColumnConfig, type ColumnGroupConfig, OnyxHeadline } from 'sit-onyx';
import { useFeedStore } from '@/stores/feed.ts';
import type { Feed } from '@/api/feeds.ts';

const feedStore = useFeedStore();

const data = computed(() => {
  if (!feedStore.feeds) {
    return [];
  }
  return feedStore.feeds.map((feed: Feed) => {
    return {
      ...feed,
      last_fetched_at: feed.last_fetched_at ? new Date(feed.last_fetched_at).toLocaleString() : '-',
    };
  });
});

const columns: ColumnConfig<Feed, ColumnGroupConfig, never>[] = [
  { key: 'name', label: 'Name' },
  { key: 'url', label: 'URL' },
  { key: 'last_fetched_at', label: 'Last Update' },
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
