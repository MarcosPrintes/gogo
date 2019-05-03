<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>APP</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <!-- <v-btn flat @click="login">
        <span class="mr-2">{{ textLoginBtn }}</span>
      </v-btn> -->
    </v-toolbar>
    <v-content>
     <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
import Sidebar from './components/Sidebar'
import axios from "axios";

export default {
  name: 'App',
  components: {
    Sidebar
  },
  data () {
    return {
    }
  },
  methods: {
  }
}
</script>
