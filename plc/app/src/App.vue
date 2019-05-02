<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>APP</span>
        <!-- <span class="font-weight-light">MATERIAL DESIGN</span> -->
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn flat @click="login">
        <span class="mr-2">Login</span>
      </v-btn>
    </v-toolbar>

    <v-content>
      <v-dialog v-model="dialog" width="300">
        <v-card>
          <!-- <v-card-title class="headline grey" primary-title>title</v-card-title> -->
           <v-card-title>   
            <span class="headline">Login</span>
          </v-card-title>
    
          <v-container grid-list-md>
            <v-layout wrap>
                <v-flex xs12 sm12 md12>
                  <v-text-field label="User Name*" v-model="user" required></v-text-field>
                </v-flex>
                <v-flex xs12 sm12 md12>
                  <v-text-field label="Password*" v-model="password" required></v-text-field>
                </v-flex>
            </v-layout>
          </v-container>
             <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click="dialog = false">Register</v-btn>
          <v-btn color="blue darken-1" flat @click="login">Login</v-btn>
        </v-card-actions>
        </v-card>
      </v-dialog>
    </v-content>
  </v-app>
</template>

<script>
// import HelloWorld from './components/HelloWorld'
import axios from "axios";

export default {
  name: 'App',
  components: {
  },
  data () {
    return {
      dialog: true,
      user:'',
      password:'',
      url: 'http://localhost:8089/api/v1'
    }
  },
  methods: {
    login() {
        // axios
        // .get(this.url+'/all')
        // .then( response => {
        //   console.log("Response => ", response);
        // })
        // .catch(err => {
        //   console.log("Error => ", err);
        // });
        axios
        .post(this.url+'/signup',{  user: 'usuario1', password: 'senhauser1' })
        .then(response =>{
          console.log(response)
        })
    },
  }
}
</script>
