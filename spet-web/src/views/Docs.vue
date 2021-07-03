<template>
  <v-app id="inspire">
    <v-app-bar
      app
      color="cyan"
      dark
    >
    </v-app-bar>

    <div class="spet-main">
        <ListDocs/>
    </div>
    <v-footer
      color="cyan"
      app
    >
    </v-footer>
  </v-app>
</template>

<script>
import ListDocs from '../components/ListDocs'

  export default {
    name: "Docs",
    components: {
      ListDocs
    },
    data: () => ({
      drawer: null,
    }),
  }
</script>

<style>
.spet-main {
  margin-top: 100px;
}
</style>