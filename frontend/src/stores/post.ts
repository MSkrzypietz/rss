import { defineStore } from 'pinia';
import type { Post, PostFilter } from '@/api/posts.ts';
import PostsAPI from '@/api/posts.ts';
import type { Feed } from '@/api/feeds.ts';
import FeedsAPI from '@/api/feeds.ts';

export const usePostStore = defineStore('post', {
  state: () => {
    return {
      posts: null as Post[] | null,
      postFilter: {
        searchText: '',
        selectedFeedIDs: [] as number[],
      } satisfies PostFilter,
    };
  },
  getters: {},
  actions: {
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
