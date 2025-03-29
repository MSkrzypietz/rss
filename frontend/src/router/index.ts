import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '@/views/LoginView.vue';
import PostsView from '@/views/PostsView.vue';
import { useAuthStore } from '@/stores/auth';
import NotFoundView from '@/views/NotFoundView.vue';
import EditView from '@/views/EditView.vue';

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean;
  }
}

export enum Routes {
  Login = 'Login',
  Posts = 'Posts',
  Edit = 'Edit',
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: Routes.Login,
      component: LoginView,
    },
    {
      path: '/',
      name: Routes.Posts,
      component: PostsView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/edit',
      name: Routes.Edit,
      component: EditView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: NotFoundView,
    },
  ],
});

router.beforeEach(async (to, from) => {
  const authStore = useAuthStore();
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return {
      name: Routes.Login,
      query: { redirect: to.fullPath },
    };
  }
});

export default router;
