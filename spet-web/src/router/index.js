import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import store from '../store/'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter(to, from, next) {
      if (to.name !== 'Login' && !store.state.user.login) next({ name: 'Login' })
      else next()
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/students',
    name: 'Students',
    component: () => import('../views/Students.vue'),
    beforeEnter(to, from, next) {
      if (to.name !== 'Login' && !store.state.user.login) next({ name: 'Login' })
      else next()
    }
  },
  {
    path: '/docs',
    name: 'Docs',
    component: () => import('../views/Docs.vue'),
    beforeEnter(to, from, next) {
      if (to.name !== 'Login' && !store.state.user.login) next({ name: 'Login' })
      else next()
    }
  },
  {
    path: '/docs/edit',
    name: 'DocsRed',
    component: () => import('../views/DocsRed.vue'),
    beforeEnter(to, from, next) {
      if (to.name !== 'Login' && !store.state.user.login) next({ name: 'Login' })
      else next()
    }
  },
  {
    path: '/comps',
    name: 'Comps',
    component: () => import('../views/Comps.vue'),
    beforeEnter(to, from, next) {
      if (to.name !== 'Login' && !store.state.user.login) next({ name: 'Login' })
      else next()
    }
  },
  {
    path: '/table',
    name: 'Table',
    component: () => import('../views/Table.vue'),
    beforeEnter(to, from, next) {
      if (to.name !== 'Login' && !store.state.user.login) next({ name: 'Login' })
      else next()
    }
  }
]

const router = new VueRouter({
    routes,
    mode: 'history'
})

export default router
