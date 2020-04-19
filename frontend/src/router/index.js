import Vue from 'vue'
import VueRouter from 'vue-router'
import store from "../store"

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    redirect: '/tasks'
  },
  {
    path: '/login',
    name: 'login',
    meta: {layout: 'empty', requiresAuth: false},
    component: () => import('../views/Login.vue')
  },
  {
    path: '/tasks',
    name: 'task',
    meta: {layout: 'main', requiresAuth: true},
    component: () => import('../views/Task.vue')
  }, 
  {
    path: '/settings',
    name: "settings",
    meta: {layout: 'main', requiresAuth: true},
    component: () => import('../views/Settings.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      next()
      return
    }
    next('/login') 
  } else {
    next() 
  }
})

export default router
