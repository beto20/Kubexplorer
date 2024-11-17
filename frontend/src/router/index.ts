import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import HomePage from "../pages/HomePage.vue";
import LoginPage from "../pages/LoginPage.vue";

const routes: Array<RouteRecordRaw> = [
    { path: '/', name: 'Home', component: HomePage },
    { path: '/login', name: 'Login', component: LoginPage },
];


const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;