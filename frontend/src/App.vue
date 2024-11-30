<script setup lang="ts">
import { RouterView } from 'vue-router';
import { OnyxAppLayout } from 'sit-onyx';
import NavBar from '@/components/NavBar.vue';
</script>

<template>
  <OnyxAppLayout>
    <NavBar />
    <RouterView />
  </OnyxAppLayout>
</template>

<style scoped></style>
