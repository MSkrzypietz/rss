import './assets/main.css';
import '@fontsource-variable/source-code-pro';
import '@fontsource-variable/source-sans-3';
import 'sit-onyx/style.css';

import { createApp } from 'vue';
import { createOnyx } from 'sit-onyx';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';

const onyx = createOnyx();
const app = createApp(App);

app.use(onyx);
app.use(createPinia());
app.use(router);

app.mount('#app');
