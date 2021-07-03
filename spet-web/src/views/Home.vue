<template>
<v-app app dark class="spet-color">
      <div class="logo">
        <v-img
          src="../assets/bigLogo.png"
          height="400px"
        ></v-img>
    </div>
      <v-layout column justify-center align-center>
          <v-col cols="12">
            <v-row
              :justify="'center'"
              style="height: 300px;"
            >
              <v-card
                  class="mx-5 scale"
                  max-width="344"
                  @click="$router.push({ path: '/students' })"
                >
                <v-img
                    src="../assets/bd.png"
                    height="200px"
                  ></v-img>

                  <v-card-title>
                    Студенты
                  </v-card-title>

                </v-card>
                <v-card
                  class="mx-5 scale"
                  max-width="344"
                  @click="$router.push({ path: '/comps' })"
                >
                  <v-img
                    src="../assets/comps.png"
                    height="200px"
                  ></v-img>

                  <v-card-title>
                    Компьютеры
                  </v-card-title>
                </v-card>
            </v-row>
          </v-col>
          <v-col cols="12">
          </v-col>
    </v-layout>
</v-app>

</template>

<script>

export default {
  name: 'Home',
}
</script>

<style scoped>
.spet-color {
    background: linear-gradient(to top, rgba(175, 255, 175, 0.7), #84bbff) !important;
}

.logo {
  position: absolute;
  width: 300px;
  margin: 50px;
}

  .scale{
    cursor: pointer;
    transition: 0.5s;
    display: block;
  }
  .scale:hover {
    transform: scale(1.1);
  }
</style>