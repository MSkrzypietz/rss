import api from '@/api/api.ts';

export type Post = {
  id: number;
  feed_id: number;
  feed_name: string;
  title: string;
  url: string;
  description: string;
  published_at: string | null;
  created_at: string;
  updated_at: string;
};

export default class PostsAPI {
  public static async getUnreadPosts(): Promise<Post[] | null> {
    const resp = await api.get<Post[]>('posts');
    if (resp.status !== 200) {
      return null;
    }
    return resp.data;
  }

  public static markPostAsRead(postID: number): Promise<void> {
    return api.post(`posts/${postID}/read`);
  }
}
