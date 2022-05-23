import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue';

const routes = [
    {
        path: '/',
        name: 'home',
        component: HomeView
    },
    {
        path: '/about',
        name: 'about',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
    },
    {
        path: '/author',
        name: 'author',
        component: () => import(/* webpackChunkName: "authors" */ '../views/AuthorView')
    },
    {
        path: '/posts',
        name: 'posts',
        component: () => import(/* webpackChunkName: "posts" */ '../views/PostsView')
    },
    {
        path: '/contact',
        name: 'contact',
        component: () => import(/* webpackChunkName: "posts" */ '../views/ContactView')
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
