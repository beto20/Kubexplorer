import {createApp} from 'vue'
import './styles/main.css';
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import router from "./router";
import App from "./App.vue";

const vuetify = createVuetify({
    components,
    directives,
});

const app = createApp(App);
console.log("router.getRoutes()", router.getRoutes());
app.use(vuetify);
app.use(router)
app.mount('#app');