<script setup lang="ts">
import { OnyxButton, OnyxHeadline, OnyxInput, OnyxPageLayout } from 'sit-onyx';
</script>

<template>
  <OnyxPageLayout>
    <div class="onyx-grid-container">
      <OnyxHeadline is="h1">Welcome Back</OnyxHeadline>
      <form class="onyx-grid">
        <OnyxInput class="onyx-grid-span-4" label="API Key" />
        <OnyxButton class="onyx-grid-span-16" label="Log In" type="submit" />
      </form>
    </div>
  </OnyxPageLayout>
</template>
