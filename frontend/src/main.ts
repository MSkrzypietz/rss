import '@fontsource-variable/source-code-pro';
import '@fontsource-variable/source-sans-3';
import 'sit-onyx/global.css';
import 'sit-onyx/style.css';

import { createApp, markRaw } from 'vue';
import { createOnyx, createToastProvider, TOAST_PROVIDER_INJECTION_KEY } from 'sit-onyx';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import type { Router } from 'vue-router';

declare module 'pinia' {
  export interface PiniaCustomProperties {
    router: Router;
  }
}

const onyx = createOnyx({ router });
const pinia = createPinia();
const app = createApp(App);

pinia.use(({ store }) => {
  store.router = markRaw(router);
});

app.use(onyx);
app.use(pinia);
app.use(router);

app.provide(TOAST_PROVIDER_INJECTION_KEY, createToastProvider());

app.mount('#app');
