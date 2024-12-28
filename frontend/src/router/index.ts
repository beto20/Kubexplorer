import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import HomePage from "../pages/HomePage.vue";
import OverviewPage from "../pages/OverviewPage.vue";
import GeneralPage from "../pages/GeneralPage.vue";
import WorkloadPage from "../pages/WorkloadPage.vue";
import NetworkPage from "../pages/NetworkPage.vue";
import StoragePage from "../pages/StoragePage.vue";
import SettingsPage from "../pages/SettingsPage.vue";
import ConnectionPage from "../pages/ConnectionPage.vue";
import DocumentationPage from "../pages/DocumentationPage.vue";

const routes: Array<RouteRecordRaw> = [
    { path: '/home', name: 'home', component: HomePage },
    { path: '/overview', name: 'overview', component: OverviewPage },
    { path: '/general', name: 'general', component: GeneralPage },
    { path: '/workload', name: 'workload', component: WorkloadPage },
    { path: '/network', name: 'network', component: NetworkPage },
    { path: '/storage', name: 'storage', component: StoragePage },
    { path: '/settings', name: 'settings', component: SettingsPage },
    { path: '/documentation', name: 'documentation', component: DocumentationPage },
    { path: '/connections', name: 'connections', component: ConnectionPage },
];


const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;