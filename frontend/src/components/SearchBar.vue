<script setup lang="ts">
import { OnyxInput, OnyxSelect, type SelectOption } from 'sit-onyx';
import { computed } from 'vue';
import { usePostStore } from '@/stores/post.ts';
import { useFeedStore } from '@/stores/feed.ts';

const postStore = usePostStore();
const feedStore = useFeedStore();

const feedOptions = computed((): SelectOption<number>[] => {
  if (feedStore.feeds === null) {
    return [];
  }

  return feedStore.feeds.map((feed) => {
    return {
      value: feed.id,
      label: feed.name,
    };
  });
});
</script>

<template>
  <div class="onyx-grid">
    <OnyxInput
      class="onyx-grid-span-8"
      v-model="postStore.postFilter.searchText"
      label="Search"
      :hideLabel="true"
      placeholder="Search"></OnyxInput>
    <OnyxSelect
      class="onyx-grid-span-8 onyx-grid-md-span-4"
      v-model="postStore.postFilter.selectedFeedIDs"
      label="Feed selection"
      :hideLabel="true"
      listLabel="Feed selection"
      :options="feedOptions"
      :withSearch="true"
      multiple
      placeholder="Feed" />
  </div>
</template>

<style scoped lang="scss"></style>
