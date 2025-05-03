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

export type PostFilter = {
  searchText: string;
  selectedFeedIDs: number[];
};

export default class PostsAPI {
  public static async getUnreadPosts(filter?: PostFilter): Promise<Post[] | null> {
    const params = this._getPostFilterQueryParams(filter);
    const resp = await api.get('posts', { params });
    if (resp.status !== 200) {
      return null;
    }
    return resp.data.posts;
  }

  private static _getPostFilterQueryParams(filter?: PostFilter): { [key: string]: any } {
    if (filter === undefined) {
      return {};
    }

    const params: { [key: string]: any } = {};

    const searchText = filter.searchText.trim();
    if (searchText !== '') {
      params['searchText'] = searchText;
    }

    const selectedFeedIDs = filter.selectedFeedIDs;
    if (selectedFeedIDs.length > 0) {
      params['feedIDs'] = selectedFeedIDs.join(',');
    }

    return params;
  }

  public static markPostAsRead(postID: number): Promise<void> {
    return api.post(`posts/${postID}/read`);
  }
}
