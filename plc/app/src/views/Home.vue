<template>
  <v-container fluid>
      <p>home</p>
      <sidebar></sidebar>
  </v-container>
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
      url: 'http://localhost:8089/api/v1',
      textLoginBtn: 'Login',
      users:[]
    }
  },
  mounted(){
    console.log('home mounted')
      axios.get(this.url+'/all')
      .then(res => {
        this.users = res.data
     })
      .catch(err =>{
        // console.log(err)
      })
  },
  methods: {
  
  }
}
</script>

<script>
  import HelloWorld from '../components/HelloWorld'

  export default {
    components: {
      HelloWorld
    }
  }
</script>
