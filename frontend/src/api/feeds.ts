import api from '@/api/api.ts';

export type Feed = {
  id: number;
  user_id: number;
  name: string;
  url: string;
  created_at: string;
  updated_at: string;
};

export default class FeedsAPI {
  public static async getFeeds(): Promise<Feed[] | null> {
    const resp = await api.get<Feed[]>('feeds');
    if (resp.status !== 200) {
      return null;
    }
    return resp.data;
  }
}
