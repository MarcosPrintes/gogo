<template>
  <div class="app-alert">
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
  </div>
</template>
<script>
export default {
}
</script>
