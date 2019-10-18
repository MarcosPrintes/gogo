<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    
  </div>
</template>

<script>

import axios from 'axios'

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },

  created() {
    axios.get('http://localhost:8082/rota')
      .then(function (response) {
        console.log('then', response);
      })
      .catch(function (error) {
        console.log('catch => ', error);
      })
      .finally(function () {
        console.log('finally get');
      });
  },

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
