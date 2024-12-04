<script setup lang="ts">
import { OnyxButton, OnyxHeadline, OnyxInput, OnyxPageLayout } from 'sit-onyx';
</script>

<template>
  <OnyxPageLayout>
    <div class="onyx-grid-container">
      <section class="header">
        <OnyxHeadline is="h1">Welcome Back</OnyxHeadline>
        <p>Enter your API key to access your RSS feed.</p>
      </section>

      <form class="onyx-grid" @submit.prevent>
        <OnyxInput class="onyx-grid-span-4" label="API Key" required />
        <OnyxButton class="onyx-grid-span-16" label="Log In" type="submit" />
      </form>
    </div>
  </OnyxPageLayout>
</template>

<style lang="scss" scoped>
.header {
  margin-bottom: var(--onyx-spacing-lg);
}
</style>
