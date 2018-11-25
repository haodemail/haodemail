import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/pages/Home'
import Dashboard from '@/pages/Dashboard'
// import Domain from '@/pages/Domain'
// import User from '@/pages/User'
// import Config from '@/pages/Config'
// import MailLog from '@/pages/MailLog'
// import Login from '@/pages/Login'

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
                }]
        }]
})
    //     {
    //       path: 'domain',
    //       name: 'Domain',
    //       component: Domain
    //     },
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
    // {
    //   path: '/login',
    //   name: 'Login',
    //   components: {
    //     blank: Login
    //   }
    // },
