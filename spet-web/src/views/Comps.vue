<template>
    <v-app id="inspire">
            <v-navigation-drawer
        v-model="drawer"
        app
        >
        <v-list dense>
             <v-card
                    v-for="item in TEMPLATES_COMPS"
                    :key="item.id"
                    class="mx-auto ma-2"
                    max-width="250"
                >

                    <v-card-text class="text--primary">
                        <div>{{item.name}}</div>
                    </v-card-text>

                    <v-card-actions>
                    <v-btn
                        color="red"
                        text
                        @click="del_templ(item)"
                    >
                        Удалить
                    </v-btn>

                    <v-btn
                        color="green"
                        text
                        @click="use_templ(item)"
                    >
                        Применить
                    </v-btn>
                    </v-card-actions>
                </v-card>
        </v-list>
        </v-navigation-drawer>
        <v-app-bar
            app
            color="green"
            dark
        >
        <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title
        style="width: 300px"
        class="ml-0 pl-4"
      >
        <span class="hidden-sm-and-down">Компьютеры</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        icon
        large
        @click="exit"
      >
        <v-avatar
          size="32px"
          item
        >
          <v-img
            src="../assets/exit.svg"
            alt="Vuetify"
          ></v-img></v-avatar>
      </v-btn>
    </v-app-bar>
    <v-container
    class="fill-height"
    fluid
    >
        <v-row
            align="center"
            justify="center"
        >
            <v-col class="text-center">
                <TableComps/>
            </v-col>
        </v-row>
    </v-container>
  </v-app>
</template>

<script>
import TableComps from "../components/TableComps";

export default {
    name: "Comps",
    components: {
        TableComps
    },
    data: () => ({
        drawer: null,
    }),
    computed: {
        TEMPLATES_COMPS() {
            return this.$store.state.students.TEMPLATES_COMPS
        },
    },
    methods: {
        del_templ({id}) {
            this.$store.commit("DELETE_TEMPLATES_COMP", id)
        },
        use_templ(i) {
            this.$store.commit("STUDENTS_COMP_TEMPL", i)
        },
        exit() {
            this.$store.dispatch("LOGOUT")
            this.$router.push({ path:'/login'})
        }
    }
}
</script>

<style>

</style>