import { defineStore } from 'pinia';
import type { Post } from '@/api/posts.ts';
import PostsAPI from '@/api/posts.ts';

export const useFeedStore = defineStore('feed', {
  state: () => {
    return {
      posts: null as Post[] | null,
    };
  },
  getters: {},
  actions: {
    async fetchUnreadPosts() {
      this.posts = await PostsAPI.getUnreadPosts();
    },
    async markPostAsRead(postID: number) {
      await PostsAPI.markPostAsRead(postID);
      if (this.posts) {
        this.posts = this.posts.filter((post: Post) => post.id !== postID);
      }
    },
  },
});
