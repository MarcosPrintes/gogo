<template>
  <v-app>
    <v-content>
      <v-dialog v-model="dialog" width="300" persistent>
        <v-card> 
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
          <v-btn color="blue darken-1" flat @click="register">Register</v-btn>
          <v-btn color="blue darken-1" flat @click="login">Login</v-btn>
        </v-card-actions>
        </v-card>
      </v-dialog>

      <v-dialog  v-model="registerDialog" persistent max-width="600px">
        <v-card>
          <v-card-title>
            <span class="headline">Cadastro de usuário</span>
          </v-card-title>
          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12>
                  <v-text-field v-model="registerName" label="Nome*" required></v-text-field>
                </v-flex>
                <v-flex xs12>
                  <v-text-field v-model="registerEmail" label="Email*" required></v-text-field>
                </v-flex>
                <v-flex xs12>
                  <v-text-field v-model="registerPassword" label="Senha*" type="password" required></v-text-field>
                </v-flex>
                <v-flex xs12 sm6>
                  <v-autocomplete v-model="registerType" :items="['Admin', 'Usuário Comun']" label="Perfil" multiple ></v-autocomplete>
                </v-flex>
              </v-layout>
            </v-container>
            <small>* Campos obrigatórios</small>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" flat @click="registerClose" >cancelar</v-btn>
            <v-btn color="blue darken-1" flat @click="save" >Salvar</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-content>
  </v-app>
</template>

<script>
import axios from "axios";

export default {
  name: 'App',
  components: {
    // Sidebar
  },
  data () {
    return {
      dialog: true,
      user:'',
      password:'',
      url: 'http://localhost:6543',
      textLoginBtn: 'Login',
      registerDialog: false,
      registerName:'',
      registerEmail:'',
      registerPassword:'',
      registerType:'',
      users:[]
    }
  },
  methods: {
    login() { 
      axios
        .post(this.url+'/login', {user_name: this.user, user_pass: this.password})
        .then(response => {
          if(response.status == 200) {
            this.dialog = false
            this.textLoginBtn = 'Logout'
            this.$store.commit('userIsLogged')
            this.$router.replace('/')
         }
        })
        .catch(err => {
          // console.log(err)
      })
    },
    register(){
      this.dialog = false
      this.registerDialog = true  
    },
    registerClose(){
      this.registerDialog = false;
      this.dialog = true;
    },
    save(){
      let payload = { 
        user_name: this.registerName, 
        user_email: this.registerEmail,
        user_pass: this.registerPassword,
        user_type: 'admin'
    }
      axios.post(this.url+'/create', payload)
      .then(response =>{
        if(response.status == 200){
          this.registerDialog = false
          this.dialog = true
          this.clean()
        }
      })
      .catch(err => {
        // console.log(err)
      })
    },
    clean(){
      this.registerName     = ''
      this.registerEmail    = ''
      this.registerPassword = ''
      this.user_type        = ''
    }
  }
}
</script>
