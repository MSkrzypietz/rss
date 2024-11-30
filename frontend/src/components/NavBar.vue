<script setup lang="ts">
import { useColorMode } from '@vueuse/core';
import { OnyxNavBar, OnyxNavButton } from 'sit-onyx';
import { OnyxColorSchemeMenuItem } from 'sit-onyx';

const { store: colorScheme } = useColorMode();
</script>

<template>
  <OnyxNavBar appName="RSS">
    <OnyxNavButton href="/feed" label="Feed" />
    <OnyxNavButton href="/edit" label="Edit" />
    <OnyxColorSchemeMenuItem v-model="colorScheme" />
  </OnyxNavBar>
</template>
