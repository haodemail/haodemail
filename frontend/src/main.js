import Vue from 'vue'
import App from './App'
import router from './router'
import iView from 'iview'
import '!style-loader!css-loader!less-loader!./theme/index.less'

Vue.config.productionTip = false

Vue.use(iView)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
