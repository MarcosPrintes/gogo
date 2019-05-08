<template>
  <v-container fluid>
    <sidebar :listUsers="users" ></sidebar>
  </v-container>
</template>

<script>
import axios from "axios";
import Sidebar from '../components/Sidebar'

export default {
  name: 'Home',
  components:{
    Sidebar
  },
  data () {
    return {
      url: 'http://localhost:6543',
      textLoginBtn: 'Login',
      users:[]
    }
  },
  mounted(){
      axios
      .get(this.url+'/get')
      .then(res =>{
        this.users = res.data
      })
      .catch(err =>{
        console.log(err)
      })
  },
}
</script>
