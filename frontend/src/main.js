import Vue from 'vue'
import App from './App'
import router from './router'
import iView from 'iview'
import '!style-loader!css-loader!less-loader!./theme/index.less'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false
Vue.use(VueAxios, axios)
Vue.use(iView)


// JWT AUTH，check TOKEN stored in localStorage
router.beforeEach(({name}, from, next) => {
  if (sessionStorage.getItem('JWT_TOKEN')) {
    // current is login page
    if (name === 'Login') {
      next('/');
    } else {
      next();
    }
  } else {
    if (name === 'Login') {
      next();
    } else {
      next({name: 'Login'});
    }
  }
});


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {App}
})
