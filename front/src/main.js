import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import Vuetify from 'vuetify'

Vue.config.productionTip = false

new Vue({
  router,
  Vuetify,  
  render: h => h(App)
}).$mount('#app')
