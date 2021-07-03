<template>
  <v-app id="inspire">
    <v-app-bar
      app
      color="cyan"
      dark
    >
    </v-app-bar>
    <iframe id="my_iframe" style="display:none;"></iframe>

    <div class="spet-main">
        <v-layout row justify-center align-center>
        <v-card max-width="600">
            <v-card-title>
            <h1>{{CURRENT_DOC_NAME}}</h1>
            </v-card-title>
            <v-card-text>
            <v-form ref="form" lazy-validation>
                <v-text-field v-for="(value, name) in CURRENT_DOC"
                    :key="value"
                    v-model="currentData[name]"
                    :label="value"
                    required
                ></v-text-field>

                    <v-btn class="green mr-4" dark @click="submit">Создать</v-btn>
            </v-form>
            </v-card-text>
        </v-card>
        </v-layout>
    </div>
    <v-footer
      color="cyan"
      app
    >
    </v-footer>
  </v-app>
</template>

<script>
  export default {
    name: "DocsRed",
    data: () => ({
      drawer: null,
      currentData: {}
    }),
    created() {
        if (!this.CURRENT_DOC_NAME) {
            this.$router.push({ path:'/docs'})
        }
    },
    computed: {
        CURRENT_DOC() {
            return this.$store.state.docs.CURRENT_DOC
        },
        CURRENT_DOC_NAME() {
            return this.$store.state.docs.CURRENT_DOC_NAME
        },
        CURRENT_DOC_RED() {
            return this.$store.state.docs.CURRENT_DOC_RED
        }
    },
    methods: {
      submit() {
        this.$store.dispatch("CREATE_DOCS_NAME", this.currentData)
      }
    }
  }
</script>

<style>
.spet-main {
  margin-top: 100px;
}
</style>