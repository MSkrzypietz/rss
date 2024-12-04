<script setup lang="ts">
import { useColorMode } from '@vueuse/core';
import { OnyxNavBar, OnyxNavButton, OnyxUserMenu, OnyxColorSchemeMenuItem } from 'sit-onyx';

const { store: colorScheme } = useColorMode();
</script>

<template>
  <OnyxNavBar class="navbar" appName="RSS">
    <OnyxNavButton href="/feed" label="Feed" />
    <OnyxNavButton href="/edit" label="Edit" />
    <OnyxUserMenu class="navbar__userMenu" username="John Doe">
      <OnyxColorSchemeMenuItem v-model="colorScheme" />
    </OnyxUserMenu>
  </OnyxNavBar>
</template>

<style lang="scss" scoped>
.navbar {
  &__userMenu {
    margin-left: auto;
  }
}
</style>
