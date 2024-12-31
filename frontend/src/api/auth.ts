import type { User } from '@/api/users';
import api from '@/api/api';

export default class AuthAPI {
  public static async login(apiKey: string): Promise<User | null> {
    const resp = await api.post<User>('auth/login', { apiKey });
    if (resp.status !== 200) {
      return null;
    }
    return resp.data;
  }
}
