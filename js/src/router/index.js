import { createRouter,createWebHistory} from "vue-router";
import {useUserStore} from "@/store/User";
const router = createRouter({
  history: createWebHistory(),
  routes: [
      {
          path: '/',
          component: () => import('@/views/layout/layoutContainer.vue'),
          redirect: '/home',
          children: [
              {path: '/home', component: () => import('@/views/home.vue')},
              {path: 'hot', name: 'hot', component: () => import('@/views/hot.vue')},
              {path: 'user', name: 'user', component: () => import('@/views/user.vue')},
              {path: 'recommend', name: 'recommend', component: () => import('@/views/recommend.vue')},
              {path: 'sport', name: 'sport', component: () => import('@/views/channel/sport.vue')},
              {path: 'entertainment', name: 'entertainment', component: () => import('@/views/channel/entertainment.vue')},
              { path: 'news', name: 'news', component: () => import('@/views/channel/news.vue') },
              {path: 'music', name: 'music', component: () => import('@/views/channel/music.vue')},
              {path: 'knowledge', name: 'knowledge', component: () => import('@/views/channel/knowledge.vue')},
              {path: 'cartoon', name: 'cartoon', component: () => import('@/views/channel/cartoon.vue')},
              {path: 'search', name: 'search', component: () => import('@/views/search.vue')},

          ]
      },
  ]})

export default router;