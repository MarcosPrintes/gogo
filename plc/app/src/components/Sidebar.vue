<template>
  <div class="sidebar">
    <p>sidebar</p>
    <!-- <div v-for="u in lusers" :key="u.id" >
      <users :name="u.name"></users>
    </div> -->
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
.side-bar{
  position: fixed;
  width: 150px;
  height: 100%;
  top: 0;
  left: 0;  
  background: rgba($color: #000000, $alpha: .5);
  padding-top: 100px;
}

</style>
