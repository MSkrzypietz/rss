import { defineStore } from 'pinia';
import type { Post, PostFilter } from '@/api/posts.ts';
import PostsAPI from '@/api/posts.ts';

export const usePostStore = defineStore('post', {
  state: () => {
    return {
      isFetchingPosts: false,
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
      try {
        this.isFetchingPosts = true;
        this.posts = await PostsAPI.getUnreadPosts(this.postFilter);
      } finally {
        this.isFetchingPosts = false;
      }
    },
    async markPostAsRead(postID: number) {
      await PostsAPI.markPostAsRead(postID);
      if (this.posts) {
        this.posts = this.posts.filter((post: Post) => post.id !== postID);
      }
    },
  },
});
