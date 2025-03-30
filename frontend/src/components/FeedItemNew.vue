<script setup lang="ts">
import { OnyxButton, OnyxHeadline, OnyxInput, OnyxToast, useToast } from 'sit-onyx';
import { onMounted, ref, useTemplateRef } from 'vue';
import FeedsAPI from '@/api/feeds.ts';
import { useFeedStore } from '@/stores/feed.ts';

const feedStore = useFeedStore();
const toast = useToast();

const formRef = useTemplateRef('form');

const newFeedName = ref('');
const newFeedUrl = ref('');

const addNewFeed = async (): Promise<void> => {
  await FeedsAPI.addNewFeed(newFeedName.value, newFeedUrl.value);
  toast.show({ headline: `Added new feed: ${newFeedName.value}` });
  newFeedName.value = '';
  newFeedUrl.value = '';
  formRef.value?.reset();
  await feedStore.fetchFeeds();
};

onMounted(async () => {
  await feedStore.fetchFeeds();
});
</script>

<template>
  <div>
    <OnyxHeadline is="h2">New Feed</OnyxHeadline>
    <form ref="form" class="onyx-grid new-feed-form" @submit.prevent="addNewFeed">
      <OnyxInput class="onyx-grid-span-4" label="Name" required v-model="newFeedName" />
      <OnyxInput class="onyx-grid-span-4" label="URL" required v-model="newFeedUrl" />
      <OnyxButton class="onyx-grid-span-16" label="Add" type="submit" />
    </form>
    <OnyxToast />
  </div>
</template>

<style scoped lang="scss">
.new-feed-form {
  margin-top: var(--onyx-spacing-lg);
}
</style>
