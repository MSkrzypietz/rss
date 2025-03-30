import { defineStore } from 'pinia';
import type { Feed } from '@/api/feeds.ts';
import FeedsAPI from '@/api/feeds.ts';

export const useFeedStore = defineStore('feed', {
  state: () => {
    return {
      feeds: null as Feed[] | null,
    };
  },
  getters: {},
  actions: {
    async fetchFeeds() {
      this.feeds = await FeedsAPI.getFeeds();
    },
  },
});
