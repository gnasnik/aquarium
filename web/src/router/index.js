import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            redirect: '/algorithm'
        },
        {
            path: '/',
            component: () => import(/* webpackChunkName: "home" */ '../components/common/Home.vue'),
            meta: { title: 'FILES' },
            children: [
                {
                    path: '/algorithm',
                    component: () => import(/* webpackChunkName: "table" */ '../components/page/Algorithm.vue'),
                    meta: { title: 'Algorithm' }
                },
                {
                    path: '/exchange',
                    component: () => import(/* webpackChunkName: "table" */ '../components/page/Exchange.vue'),
                    meta: { title: 'Exchange' }
                },
                {
                    path: '/userinfo',
                    component: () => import(/* webpackChunkName: "table" */ '../components/page/User.vue'),
                    meta: { title: 'User' }
                },
                {
                    // 富文本编辑器组件
                    path: '/editor',
                    name: 'editor',
                    component: () => import(/* webpackChunkName: "editor" */ '../components/page/VueEditor.vue'),
                    meta: { title: 'Editor' }
                },
                {
                    path: '/404',
                    component: () => import(/* webpackChunkName: "404" */ '../components/page/404.vue'),
                    meta: { title: '404' }
                },
                {
                    path: '/403',
                    component: () => import(/* webpackChunkName: "403" */ '../components/page/403.vue'),
                    meta: { title: '403' }
                }
            ]
        },
        {
            path: '/login',
            component: () => import(/* webpackChunkName: "login" */ '../components/page/Login.vue'),
            meta: { title: 'Login' }
        },
        {
            path: '*',
            redirect: '/404'
        }
    ]
});