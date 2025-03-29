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

export enum Route {
  Login = 'Login',
  Posts = 'Posts',
  Edit = 'Edit',
}

export const RoutePath: { [route in Route]: string } = {
  [Route.Login]: '/login',
  [Route.Posts]: '/',
  [Route.Edit]: '/edit',
};

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: RoutePath.Login,
      name: Route.Login,
      component: LoginView,
    },
    {
      path: RoutePath.Posts,
      name: Route.Posts,
      component: PostsView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: RoutePath.Edit,
      name: Route.Edit,
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
      name: Route.Login,
      query: { redirect: to.fullPath },
    };
  }
});

export default router;
