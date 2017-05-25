import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import login from '@/components/login'
import manage from '@/components/manage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/login',
      name: 'Scan to Login',
      component: login
    },
    {
      path: '/manage',
      name: 'Manage',
      component: manage
    }
  ]
})
