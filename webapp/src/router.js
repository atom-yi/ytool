import { createRouter, createWebHashHistory } from 'vue-router';

const welcome = () => import('./views/Welcome.vue');
const jsonView = () => import('./views/Json.vue');
const httpView = () => import('./views/Http.vue');

const routes = [
    {
        path: '/',
        redirect: '/welcome'
    },
    {
        path: '/welcome',
        component: welcome,
    },
    {
        path: '/json',
        component: jsonView,
    },
    {
        path: '/http',
        component: httpView,
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router
