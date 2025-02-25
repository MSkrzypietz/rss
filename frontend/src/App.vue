<script setup lang="ts">
import { RouterView } from 'vue-router';
import { OnyxAppLayout } from 'sit-onyx';
import NavBar from '@/components/NavBar.vue';
</script>

<template>
  <OnyxAppLayout class="onyx-grid-max-lg onyx-grid-center">
    <template #navBar>
      <NavBar />
    </template>
    <RouterView />
  </OnyxAppLayout>
</template>

<style scoped></style>
