import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import vuelidate from 'vuelidate'

import 'materialize-css/dist/js/materialize.min'

Vue.config.productionTip = false

const user = JSON.parse(localStorage.getItem('user_info_3d_mesh'))
if (user) {
  axios.defaults.headers.common['Authorization'] = "Bearer_" + user.session_token
}

Vue.use(vuelidate)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
