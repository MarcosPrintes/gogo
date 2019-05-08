<template>
  <div>
    <div class="nucashers-component">
      <div class="image-user">
        <img :src="image" alt="">
      </div>
      <span>{{ name }}</span>
    </div>
  </div>
</template>
<script>
export default {
  name:'users',
 data(){
   return{
      image:'https://www.viawater.nl/files/default-user.png'
   }
 },
 props:{
   name:{
     type: String,
     required: false
   }
 } 
}
</script>
<style lang="scss" scoped>
.nucashers-component{
  margin: 0 auto 15px;
  text-align: center;
  .image-user{
    width: 60px;
    height: 60px;
    box-shadow: 3px 3px 3px #ccc;
    border-radius: 100px;
    margin: 0 auto;
    position: relative;
    overflow: hidden;
    transition: all .3s linear;

     &:hover{
      cursor: pointer;
    }
    img{
      position: absolute;
      top:50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 100%;
    }
  }
  span{
    margin-top: 5px;
    display: inline-block;
    color:#000;
  }
}
</style>
