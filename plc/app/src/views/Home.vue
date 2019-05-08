<template>
  <v-container fluid>
      <p>home</p>
      <sidebar></sidebar>        
  </v-container>
</template>

<script>
import axios from "axios";
import Sidebar from '../components/Sidebar'

export default {
  name: 'Home',
  components: {
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
    console.log('sdfsdf')
  },
  created(){
    console.log('sdfsdf cre')
  },
}
</script>

<script>

  export default {
    components: {
    
    }
  }
</script>
