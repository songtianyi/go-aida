import Home from '@/components/HelloFromVux'
import Login from '@/components/login'
import Manage from '@/components/manage'

export const routes = [
  {
    path: '/',
    component: Home
  },
  {
    path: '/login',
    component: Login
  },
  {
    path: '/manage',
    component: Manage
  }
]
