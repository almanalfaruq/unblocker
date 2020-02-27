<template>
  <v-app id="main-app">
    <v-content>
      <v-container fluid> 
        <div class="d-flex flex-column mb-6">
          <v-text-field label="Link Website" placeholder="www.contoh.com" v-model="url" />
          <v-btn color="primary" @click="writeToHosts">Unblock Web!</v-btn>
          <h4 class="font-weight-bold text-center red--text mt-4">{{ formattedError }}</h4>
        </div> 
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  props: {
    source: String
  },
  data: () => ({
    url: '',
    errorMessage: '',
  }),
  computed: {
    formattedError() {
      return this.errorMessage.replace(/^\w/, c => c.toUpperCase());
    },
  },
  methods: {
    writeToHosts() {
      window.backend.writeToHosts(this.url).then(res => {
        this.errorMessage = res
      });
    },
  },
};
</script>

<style>
</style>