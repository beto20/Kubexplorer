import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

const routes: Array<RouteRecordRaw> = [
    { path: '/home', name: 'home', component: () => import('./pages/HomePage.vue') },
    { path: '/overview', name: 'overview', component: () => import('./pages/OverviewPage.vue') },
    {
        path: '/general',
        name: 'general',
        component: () => import('./pages/general/GeneralPage.vue'),
        children: [
            {
                path: 'node',
                name: 'node',
                component: () => import('./pages/general/NodesPage.vue')
            },
            {
                path: 'namespace',
                name: 'namespace',
                component: () => import('./pages/general/NamespacePage.vue')
            },
            {
                path: 'event',
                name: 'event',
                component: () => import('./pages/general/EventsPage.vue')
            },
        ]
    },
    {
        path: '/workload',
        name: 'workload',
        component: () => import('./pages/workloads/WorkloadPage.vue'),
        children: [
            { path: 'pod', name: 'pod', component: () => import('./pages/workloads/PodPage.vue') },
            { path: 'deployment', name: 'deployment', component: () => import('./pages/workloads/DeploymentPage.vue') },
        ]
    },
    { path: '/network', name: 'network', component: () => import('./pages/network/NetworkPage.vue') },
    { path: '/storage', name: 'storage', component: () => import('./pages/storage/StoragePage.vue') },
    { path: '/settings', name: 'settings', component: () => import('./pages/common/SettingsPage.vue') },
    { path: '/documentation', name: 'documentation', component: () => import('./pages/common/DocumentationPage.vue') },
    { path: '/connections', name: 'connections', component: () => import('./pages/common/ConnectionPage.vue') },


];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;