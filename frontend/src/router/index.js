import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/pages/Home'
import Login from '@/pages/Login'
import Dashboard from '@/pages/Dashboard'
import Account from '@/pages/Account'
import Domain from '@/pages/Domain'
// import User from '@/pages/User'
// import Config from '@/pages/Config'
// import MailLog from '@/pages/MailLog'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: Home,
      children: [
        {
          path: '/',
          name: 'Dashboard',
          component: Dashboard
        }, {
          path: 'domain',
          name: 'Domain',
          component: Domain
        }, {
          path: 'account/:domain/:ID',
          name: 'Account',
          component: Account
        },
      ]
    }, {
      path: '/login',
      name: 'Login',
      component: Login
    },
    ]
})

//     {
//       path: 'user',
//       name: 'User',
//       component: User
//     },
//     {
//       path: 'config',
//       name: 'config',
//       component: Config
//     },
//     {
//       path: 'maillog',
//       name: 'maillog',
//       component: MailLog
//     },
//   ]
// },

