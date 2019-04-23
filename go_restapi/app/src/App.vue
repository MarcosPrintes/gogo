<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>GOLANG</span>
        <span class="font-weight-light">APP</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <v-content>
      <v-container grid-list-md>
        <v-layout row wrap>
          <v-flex v-for="m in movies" :key="m.name" xs6 lg3 class="m-card">
            <v-card dark color="secondary" @click="showMovieById(m.id)">
              <v-img :src="m.thumb_image" aspect-ratio="2.75"></v-img>
              <v-card-title primary-title>
                <p class>{{ m.name }}</p>
                <div>{{ m.description }}</div>
              </v-card-title>
              <v-card-text class="px-0"></v-card-text>
            </v-card>
          </v-flex>
        </v-layout>
        <v-btn color="info" @click="getAll">GET ALL</v-btn>
        <v-btn color="warning" @click="deleteAll" >DELETE</v-btn>
      </v-container>
    </v-content>

    <v-dialog v-model="dialog" width="500">
      <v-card>
        <v-card-title class="headline grey lighten-2" primary-title>{{ movieById.name }}</v-card-title>
        <v-card-text>{{ movieById.description }}</v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" flat @click="dialog = false">close</v-btn>
          <v-btn color="primary" flat @click="deleteById(movieById.id)">delete</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-app>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
  data() {
    return {
      movies: [],
      dialog: false,
      movieById:{}
    };
  },
  mounted() {},
  methods: {
    getAll() {
      axios
        .get("http://localhost:3002/api/v1/movies")
        .then(response => {
          this.movies = response.data;
          console.log(this.movies);
        })
        .catch(err => {
          console.log("Error => ", err);
        });
    },
    showMovieById(id) {
      this.dialog = true
      axios
        .get(`http://localhost:3002/api/v1/movies/${id}`)
        .then(res => {

          this.movieById = res.data
        })
        .catch(err => {});
    },
    deleteAll(){
      this.movies = []
    },
    deleteById(id){
      this.dialog = false
      axios.delete(`http://localhost:3002/api/v1/movies/delete/${id}`, {id:id})
      .then(res =>{
        console.log("delete res",res)
      }).catch(err =>{
        console.log("delete error", error)
      })

    }
  }
};
</script>
<style lang="scss" scoped>
.m-card{
  cursor: pointer;
}
</style>

