// import 'vite/modulepreload_polyfill'
import './assets/style.css'
import 'virtual:windi.css'
import { createApp } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'
import { VueQueryPlugin } from 'vue-query'
import Notifications from 'notiwind'

import App from './App.vue';

// router
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      name: 'home',
      path: '/',
      component: () => import('./pages/Home.vue')
    },
    {
      name: 'post-page',
      path: '/posts/:post_id',
      component: () => import('./pages/Post.vue')
    }
  ]
})

// app
createApp(App)
  .use(router)
  .use(Notifications)
  .use(VueQueryPlugin)
  .mount('#app');