<script setup lang="ts">
import { useColorMode } from '@vueuse/core';
import { OnyxNavBar, OnyxNavButton, OnyxUserMenu, OnyxColorSchemeMenuItem } from 'sit-onyx';
import { useAuthStore } from '@/stores/auth';
import { Route, RoutePath } from '@/router';
import { useRouter } from 'vue-router';

const { store: colorScheme } = useColorMode();
const authStore = useAuthStore();
const router = useRouter();
</script>

<template>
  <OnyxNavBar class="navbar" appName="RSS" @navigateToStart="router.push({ name: Route.Posts })">
    <template v-if="authStore.user">
      <OnyxNavButton :link="RoutePath.Posts" label="Posts" />
      <OnyxNavButton :link="RoutePath.Feeds" label="Feeds" />
      <OnyxUserMenu class="navbar__userMenu" :fullName="authStore.user.name">
        <OnyxColorSchemeMenuItem v-model="colorScheme" />
      </OnyxUserMenu>
    </template>
  </OnyxNavBar>
</template>

<style lang="scss" scoped>
.navbar {
  &__userMenu {
    margin-left: auto;
  }
}
</style>
