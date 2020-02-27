<template>
  <v-app id="main-app">
    <v-content>
      <v-container fluid> 
        <div class="d-flex flex-column mb-6">
          <v-text-field label="Link Website" placeholder="www.contoh.com" v-model="url" />
          <v-btn color="primary" @click="writeToHosts">Unblock Web!</v-btn>
          <h4 class="font-weight-bold text-center green--text mt-4" v-if="success">Success</h4>
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
    success: false,
  }),
  computed: {
    formattedError() {
      return this.errorMessage.replace(/^\w/, c => c.toUpperCase());
    },
  },
  methods: {
    writeToHosts() {
      this.errorMessage = '';
      window.backend.writeToHosts(this.url).then(res => {
        this.errorMessage = res;
        if (this.errorMessage !== '') {
          this.success = false;
        } else {
          this.success = true; 
        }
      });
    },
  },
};
</script>

<style>
</style>