<template>
  <v-container>
        <v-card>
          <v-card-title>
          <v-col cols="12">
            <v-row>
              <v-col>
                <v-col class="pa-6">
                  <v-row>
                    <v-checkbox v-model="heders.id"></v-checkbox>
                    <v-text-field v-model="filtres.id" label="ID"></v-text-field>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.surname"></v-checkbox>
                    <v-text-field v-model="filtres.surname" label="Фамилия"></v-text-field>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.name"></v-checkbox>
                    <v-text-field v-model="filtres.name" label="Имя"></v-text-field>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.middleName"></v-checkbox>
                    <v-text-field v-model="filtres.middleName" label="Отчество"></v-text-field>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.login"></v-checkbox>
                    <v-text-field v-model="filtres.login" label="Логин"></v-text-field>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.password"></v-checkbox>
                    <v-text-field v-model="filtres.password" label="Пароль"></v-text-field>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.activs"></v-checkbox>
                    <v-text-field v-model="filtres.activs" label="Активность"></v-text-field>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.numGroup"></v-checkbox>
                    <v-select
                      v-model="filtres.numGroup"
                      :items="GROUP_LIST"
                      attach
                      chips
                      label="Группа"
                      multiple
                    ></v-select>
                  </v-row>
                  <v-row>
                    <v-checkbox v-model="heders.gender"></v-checkbox>
                    <v-select
                      v-model="filtres.gender"
                      :items="ITEMS.gender"
                      attach
                      chips
                      label="Пол"
                      multiple
                    ></v-select>
                  </v-row>
                </v-col>
              </v-col>
              <v-col>
                <v-col
                  class="pa-6"
                  outlined
                  tile
                >
                  <v-row>
                    <v-col class="pa-1">
                      <v-row>
                        <v-checkbox v-model="heders.date_b"></v-checkbox>
                        <v-text-field v-model="filtres.date_b" label="Дата рождения"></v-text-field>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.city"></v-checkbox>
                        <v-text-field v-model="filtres.city" label="Нас. пункт"></v-text-field>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.street"></v-checkbox>
                        <v-text-field v-model="filtres.street" label="Улица"></v-text-field>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.house"></v-checkbox>
                        <v-text-field v-model="filtres.house" label="Дом"></v-text-field>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.flat"></v-checkbox>
                        <v-text-field v-model="filtres.flat" label="Квартира"></v-text-field>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.phone"></v-checkbox>
                        <v-text-field v-model="filtres.phone" label="Домашний телефон"></v-text-field>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.status"></v-checkbox>
                        <v-select
                          v-model="filtres.status"
                          :items="ITEMS.status"
                          attach
                          chips
                          label="Статус"
                          multiple
                        ></v-select>
                      </v-row>
                      <v-row>
                      <v-checkbox v-model="heders.budget"></v-checkbox>
                      <v-select
                        v-model="filtres.budget"
                        :items="ITEMS.budget"
                        attach
                        chips
                        label="Основа"
                        multiple
                      ></v-select>
                    </v-row>
                      <v-row>
                        <v-row
                          align="center"
                          justify="center"
                        >
                          <v-row justify="center"
                                  class="pa-2"
                          >
                            <v-btn
                              @click="set_filtrs"

                              color="blue"
                              dark
                            >
                            Применить фильтры
                            </v-btn>
                          </v-row>
                          <v-row
                          class="pa-2"
                          justify="center">
                            <v-btn
                              @click="OPEN_TEMPL"
                              color="green"
                              dark
                            >
                            Сохранить шаблон
                            </v-btn>
                          </v-row>
                          <v-row
                          class="pa-2"
                          justify="center">
                            <v-btn
                              @click="dis_diltrs"
                              color="red"
                              dark
                            >
                            Сбросить фильтры
                            </v-btn>
                          </v-row>
                        </v-row>
                      </v-row>
                    </v-col>
                  </v-row>
                </v-col>
              </v-col>
              <v-col>
                <v-col
                  class="pa-6"
                  outlined
                  tile
                >
                <v-col class="pa-1">
                    <v-row class="pa-2">
                      <v-select
                        v-model="dop"
                        :items="ITEMS.dop"
                        attach
                        chips
                        label="Дополнительное"
                        multiple
                      ></v-select>
                    </v-row>
                      <v-row>
                          <v-checkbox v-model="heders.orphan" :label="'Сирота'"></v-checkbox>
                      </v-row>
                      <v-row>
                          <v-checkbox v-model="heders.invalid" :label="'Инвалид'"></v-checkbox>
                      </v-row>
                      <v-row>
                          <v-checkbox v-model="heders.low_income" :label="'Мал. обес.'"></v-checkbox>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.idn" :label="'ИДН'"></v-checkbox>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.kdn" :label="'КДН'"></v-checkbox>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.many_children" :label="'Многодет.'"></v-checkbox>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.low_parents" :label="'Непол. семья'"></v-checkbox>
                      </v-row>
                      <v-row>
                        <v-checkbox v-model="heders.info" :label="'Доп. информация'"></v-checkbox>
                      </v-row>
                    </v-col>
                </v-col>
              </v-col>
            </v-row>
          </v-col>
            <v-row>
              <v-col>
                <v-text-field
                  v-model="search"
                  append-icon="mdi-magnify"
                  label="Быстрый поиск"
                  single-line
                  hide-details
                ></v-text-field>
              </v-col>
            </v-row>
          </v-card-title>
      <v-data-table
          :loading = LOADING
          loading-text="Загрузка"
          v-model="SELECTED"
          :headers="TABLE_HEADERS"
          :items="SORT_STUDENTS"
          :single-select="singleSelect"
          :search="search"
          item-key="id"
          show-select
          class="elevation-1"
      >
          <template v-slot:top>
          <v-switch v-model="singleSelect" label="Одиночный выбор" class="pa-3"></v-switch>
          </template>
          <template v-slot:item.orphan="{ item }">
            <v-simple-checkbox v-model="item.orphan" disabled></v-simple-checkbox>
          </template>
          <template v-slot:item.invalid="{ item }">
            <v-simple-checkbox v-model="item.invalid" disabled></v-simple-checkbox>
          </template>
          <template v-slot:item.low_income="{ item }">
            <v-simple-checkbox v-model="item.low_income" disabled></v-simple-checkbox>
          </template>
          <template v-slot:item.low_parents="{ item }">
            <v-simple-checkbox v-model="item.low_parents" disabled></v-simple-checkbox>
          </template>
          <template v-slot:item.idn="{ item }">
            <v-simple-checkbox v-model="item.idn" disabled></v-simple-checkbox>
          </template>
          <template v-slot:item.kdn="{ item }">
            <v-simple-checkbox v-model="item.kdn" disabled></v-simple-checkbox>
          </template>
          <template v-slot:item.many_children="{ item }">
            <v-simple-checkbox v-model="item.many_children" disabled></v-simple-checkbox>
          </template>
      </v-data-table>
      </v-card>
          <v-dialog v-model="dialog_templ" persistent max-width="600px">
            <v-card>
                <v-card-title>
                <span class="headline">Задайте имя шаблону фильтров</span>
                </v-card-title>
                <v-card-text>
                <v-container>
                    <v-text-field v-model="temple.name" label="Название" required></v-text-field>
                </v-container>
                </v-card-text>
                <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="red darken-1" text @click="dialog_templ = false">Отмена</v-btn>
                <v-btn color="blue darken-1" text @click="SAVE_TEMPL">Сохранить</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
  </v-container>
</template>

<script>
export default {
    name: "TableStudents",
    data () {
    return {
        search: '',
        dialog_templ: false,
        singleSelect: true,
        temple: {
          id: null,
          name: "",
          dop: [],
          headers: {},
          filtres: {}
        },
        heders: {
          id: false,
          surname: true,
          name: true,
          middleName: true,
          date_b: false,
          city: false,
          street: false,
          house: false,
          flat: false,
          phone: false,
          activs: false,
          info: false,
          gender: false,
          numGroup: true,
          status: false,
          login: false,
          password: false,
          budget: false,
          orphan: false,
          invalid: false,
          low_income: false,
          low_parents: false,
          idn: false,
          kdn: false,
          many_children: false,
        },
        dop: [],
        filtres:
        {
          id: "",
          surname: "",
          name: "",
          middleName: "",
          date_b: "",
          city: "",
          street: "",
          house: "",
          flat: "",
          phone: "",
          info: "",
          activs: "",
          gender: "",
          numGroup: "",
          status: "",
          login: "",
          password: "",
          budget: "",
          orphan: false,
          invalid: false,
          low_income: false,
          low_parents: false,
          idn: false,
          kdn: false,
          many_children: false,
        }
    }
    },
    computed: {
        FILTRES() {
            return this.$store.state.students.FILTRES
        },
        HEADERS() {
            return this.$store.state.students.HEADERS
        },
        SORT_STUDENTS() {
            return this.$store.state.students.SORT_STUDENTS
        },
        GROUP_LIST() {
            return this.$store.state.students.GROUP_LIST
        },
        LOADING() {
            return this.$store.state.students.LOADING
        },
        TABLE_HEADERS() {
            return this.$store.state.students.TABLE_HEADERS
        },
        ITEMS() {
            return this.$store.state.students.ITEMS
        },
        SELECTED: {
          get() {
              return this.$store.state.students.SELECTED
          },
          set(val) {
              this.$store.commit("UPDATE_SELECTED", val)
          }
        }
    },
    created() {
      this.$store.dispatch("GET_STUDENTS")
    },
    methods: {
      set_filtrs() {
        if (Object.prototype.hasOwnProperty.call(this.dop, "key")) {
          this.filtres.orphan = false;
          this.filtres.invalid = false;
          this.filtres.low_income = false;
          this.filtres.many_children = false;
          this.filtres.idn = false;
          this.filtres.low_parents = false;
        }
        for(let i in this.dop) {
          if (this.dop[i] == "Сирота") {
            this.filtres.orphan = true;
          } else {
            this.filtres.orphan = false;
          }
          if (this.dop[i] == "Инвалид") {
            this.filtres.invalid = true;
          } else {
            this.filtres.invalid = false;
          }
          if (this.dop[i] == "Мал. Обес.") {
            this.filtres.low_income = true;
          } else {
            this.filtres.low_income = false;
          }
          if (this.dop[i] == "Многодет.") {
            this.filtres.many_children = true;
          } else {
            this.filtres.many_children = false;
          }
          if (this.dop[i] == "ИДН") {
            this.filtres.idn = true;
          } else {
            this.filtres.idn = false;
          }
          if (this.dop[i] == "КДН") {
            this.filtres.kdn = true;
          } else {
            this.filtres.kdn = false;
          }
          if (this.dop[i] == "Непол. семья") {
            this.filtres.low_parents = true;
          } else {
            this.filtres.low_parents = false;
          }
        }
        this.$store.commit("UPDATE_HEADERS_TABLES", {
          headers: this.heders,
          filtres: this.filtres
        })
        this.$store.commit("UPDATE_FILTRES", {
          filtres: this.filtres,
        })
        this.$store.commit("UPDATE_HEADERS", {
          headers: this.heders,
        })
      },
      OPEN_TEMPL() {
        this.dialog_templ = true;
      },
      SAVE_TEMPL() {
        this.$store.commit("UPDATE_TEMPLATES", {
          id: 0,
          name: this.temple.name,
          dop: this.dop,
          headers: this.heders,
          filtres: this.filtres
        })
        this.temple = {
          id: null,
          name: "",
          dop: [],
          headers: {},
          filtres: {}
        },
        this.dialog_templ = false;
      },
      dis_diltrs() {
        this.heders = {
          id: false,
          surname: true,
          name: true,
          middleName: true,
          date_b: false,
          city: false,
          street: false,
          house: false,
          flat: false,
          phone: false,
          activs: false,
          info: false,
          gender: false,
          numGroup: true,
          status: false,
          login: false,
          password: false,
          budget: false,
          orphan: false,
          invalid: false,
          low_income: false,
          low_parents: false,
          idn: false,
          kdn: false,
          many_children: false,
        },
        this.filtres = {
          id: "",
          surname: "",
          name: "",
          middleName: "",
          date_b: "",
          city: "",
          street: "",
          house: "",
          flat: "",
          phone: "",
          info: "",
          activs: "",
          gender: "",
          numGroup: "",
          status: "",
          login: "",
          password: "",
          budget: "",
          orphan: false,
          invalid: false,
          low_income: false,
          low_parents: false,
          idn: false,
          kdn: false,
          many_children: false,
        }
        this.dop = [];
        this.$store.commit("UPDATE_HEADERS_TABLES", {
          headers: this.heders,
          filtres: this.filtres
        })
      }
    }
}
</script>

<style>

.spet_grid {
  display: flex;
  flex-wrap: wrap
}

</style>