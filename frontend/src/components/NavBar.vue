<script setup lang="ts">
import { useColorMode } from '@vueuse/core';
import { OnyxNavBar, OnyxNavButton, OnyxUserMenu, OnyxColorSchemeMenuItem } from 'sit-onyx';
import { useAuthStore } from '@/stores/auth';
import { Routes } from '@/router';
import { useRouter } from 'vue-router';

const { store: colorScheme } = useColorMode();
const authStore = useAuthStore();
const router = useRouter();
</script>

<template>
  <OnyxNavBar class="navbar" appName="RSS" @navigateToStart="router.push({ name: Routes.Posts })">
    <template v-if="authStore.user">
      <RouterLink :to="{ name: Routes.Posts }"><OnyxNavButton href="#" label="Posts" /></RouterLink>
      <RouterLink :to="{ name: Routes.Edit }"><OnyxNavButton href="#" label="Edit" /></RouterLink>
      <OnyxUserMenu class="navbar__userMenu" :username="authStore.user.name">
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
