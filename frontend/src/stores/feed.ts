import { defineStore } from 'pinia';
import type { Post, PostFilter } from '@/api/posts.ts';
import PostsAPI from '@/api/posts.ts';
import type { Feed } from '@/api/feeds.ts';
import FeedsAPI from '@/api/feeds.ts';

export const useFeedStore = defineStore('feed', {
  state: () => {
    return {
      posts: null as Post[] | null,
      feeds: null as Feed[] | null,
      postFilter: {
        searchText: '',
        selectedFeedIDs: [] as number[],
      } satisfies PostFilter,
    };
  },
  getters: {},
  actions: {
    async fetchFeeds() {
      this.feeds = await FeedsAPI.getFeeds();
    },
    async fetchUnreadPosts() {
      this.posts = await PostsAPI.getUnreadPosts(this.postFilter);
    },
    async markPostAsRead(postID: number) {
      await PostsAPI.markPostAsRead(postID);
      if (this.posts) {
        this.posts = this.posts.filter((post: Post) => post.id !== postID);
      }
    },
  },
});
