import api from '@/api/api';

export interface User {
  name: string;
  apiKey: string;
}

export default class UsersAPI {
  public static async getCurrentUser(): Promise<User | null> {
    const resp = await api.get<User>('users');
    if (resp.status !== 200) {
      return null;
    }
    return resp.data;
  }
}
