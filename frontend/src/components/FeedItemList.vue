<script setup lang="ts">
import { computed, h } from 'vue';
import sync from '@sit-onyx/icons/sync.svg?raw';
import {
  OnyxDataGrid,
  type ColumnConfig,
  type ColumnGroupConfig,
  OnyxHeadline,
  type TypeRenderMap,
  createFeature,
  OnyxSystemButton,
} from 'sit-onyx';
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

type ActionButtonsType = 'actionButtons';

const columns: ColumnConfig<Feed, ColumnGroupConfig, ActionButtonsType>[] = [
  { key: 'name', label: 'Name' },
  { key: 'url', label: 'URL' },
  { key: 'last_fetched_at', label: 'Last Update' },
  { key: 'id', label: '', type: 'actionButtons' },
];

const withActionButtonsType = createFeature(() => ({
  name: Symbol('action buttons'),
  typeRenderer: {
    actionButtons: {
      cell: {
        tdAttributes: {
          style: { width: '1.5rem' },
        },
        component: (props) => {
          return h(
            'div',
            {
              style: {
                display: 'flex',
                flexDirection: 'row',
                gap: 'var(--onyx-spacing-xs)',
              },
            },
            [
              h(OnyxSystemButton, {
                label: 'Update feed',
                icon: sync,
                onClick: () => feedStore.updateFeed(props.row.id),
              }),
            ],
          );
        },
      },
    },
  } satisfies TypeRenderMap<Feed, ActionButtonsType>,
}));

const features = [withActionButtonsType()];
</script>

<template>
  <div>
    <OnyxHeadline is="h2">All Feeds</OnyxHeadline>
    <OnyxDataGrid :columns :data :features />
  </div>
</template>

<style scoped lang="scss">
.onyx-data-grid {
  margin-top: var(--onyx-spacing-lg);
}
</style>
