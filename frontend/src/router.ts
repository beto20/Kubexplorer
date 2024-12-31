import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import HomePage from "../pages/HomePage.vue";
import OverviewPage from "../pages/OverviewPage.vue";
import GeneralPage from "../GeneralPage.vue";
import WorkloadPage from "../pages/workloads/WorkloadPage.vue";
import NetworkPage from "../pages/network/NetworkPage.vue";
import StoragePage from "../pages/storage/StoragePage.vue";
import SettingsPage from "../pages/common/SettingsPage.vue";
import ConnectionPage from "../pages/common/ConnectionPage.vue";
import DocumentationPage from "../pages/common/DocumentationPage.vue";
import PodPage from "../pages/workloads/PodPage.vue";
import DeploymentPage from "../pages/workloads/DeploymentPage.vue";

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

    { path: '/pod', name: 'pod', component: PodPage },
    { path: '/deployment', name: 'deployment', component: DeploymentPage },
];


const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;