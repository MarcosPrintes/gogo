<template>
  <div class="sidebar">
    <div v-for="u in lusers" :key="u.id" >
      <users :name="u.user_name"></users>
    </div>
  </div>
</template>
<script>
import Users from '@/components/Users'

export default {
  name: 'sidebar',
 
  data(){
    return {
    lusers: this.listUsers
   }
  },
  components: {
    Users
  },
  props:{
    listUsers:{
      type: Array
    }
  },
  watch: {
    listUsers(){
      this.lusers = this.listUsers
    }
  }
}
</script>
<style lang="scss" scoped>
.sidebar{
  position: fixed;
  width: 150px;
  height: 100%;
  top: 0;
  left: 0;  
  background: #fff;
  box-shadow: 3px 5px 7px #ccc;
  padding-top: 100px;
}

</style>
