import api from '@/api/api.ts';

export type Feed = {
  id: number;
  user_id: number;
  name: string;
  url: string;
  last_fetched_at: string | null;
  created_at: string;
  updated_at: string;
};

export default class FeedsAPI {
  public static async getFeeds(): Promise<Feed[] | null> {
    const resp = await api.get('feeds');
    if (resp.status !== 200) {
      return null;
    }
    return resp.data.feeds;
  }

  public static async addNewFeed(name: string, url: string): Promise<void> {
    await api.post('feeds', { name, url });
  }

  public static async updateFeed(id: number): Promise<void> {
    await api.post(`feeds/${id}/fetch`, {});
  }
}
