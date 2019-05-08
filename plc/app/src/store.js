import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    logged: false
  },
  mutations: {
    userIsLogged(state) {
      state.logged = true
    }
  },
  actions: {
  }
})
