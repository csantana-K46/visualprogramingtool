import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/programs',
    name: 'programs',
    component: () => import('../views/ProgramList.vue')
  },
  {
    path: '/program/:id',
    name: 'program-details',
    component: () => import('../views/Program.vue')
  },
  {
    path: '/add',
    name: 'add',
    component: () => import('../views/Program.vue')
  },
]

const router = new VueRouter({
  routes
})

export default router
