import { defineStore } from 'pinia';
import type { User } from '@/api/users';
import AuthAPI from '@/api/auth';
import UsersAPI from '@/api/users';
import { Route } from '@/router';
import type { Router } from 'vue-router';

function redirectUserAfterLogin(router: Router) {
  const { query } = router.currentRoute.value;
  if (typeof query.redirect === 'string') {
    router.push(query.redirect);
  } else {
    router.push({ name: Route.Posts });
  }
}

export const useAuthStore = defineStore('auth', {
  state: () => {
    return {
      user: null as User | null,
      isLoggingIn: false,
    };
  },
  getters: {
    isAuthenticated: (state) => {
      return state.user !== null;
    },
  },
  actions: {
    async login(apiKey: string) {
      try {
        this.isLoggingIn = true;
        this.user = await AuthAPI.login(apiKey);
        if (this.user) {
          redirectUserAfterLogin(this.router);
        }
      } finally {
        this.isLoggingIn = false;
      }
    },
    async tryAutoLogin() {
      try {
        this.isLoggingIn = true;
        this.user = await UsersAPI.getCurrentUser();
        if (this.user) {
          redirectUserAfterLogin(this.router);
        }
      } finally {
        this.isLoggingIn = false;
      }
    },
  },
});
