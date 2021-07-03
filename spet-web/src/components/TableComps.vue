<template>
    <v-container>
        <v-card
        style="min-height: 80vh"
        >
            <v-card-title>
            </v-card-title>
            <v-card-text
                        align="center"
                        justify="center"
            >
                <v-select
                style="max-width: 300px"
                :items="ROOM_LIST"
                label="Кабинет"
                v-model="room"
                ></v-select>
                <v-select
                style="max-width: 300px"
                :items="GROUP_LIST"
                label="Группа"
                v-model="group"
                @change="selectGroup"
                ></v-select>
                <v-btn
                @click="UseAll()"
                style="margin-right: 10px; margin-bottom: 30px;" color="success">Подключить всех</v-btn>
                <v-btn
                @click="SaveTempl()"
                style="margin-right: 10px; margin-bottom: 30px;" color="primary">Сохранить шаблон</v-btn>
                <v-btn
                @click="QuitAll()"
                style="margin-right: 10px; margin-bottom: 30px;" color="error">Отключить всех</v-btn>
                <v-simple-table dense style="max-width: 1200px">
                    <template v-slot:default>
                        <thead>
                            <tr>
                                <th class="text-center">Id</th>
                                <th class="text-center">Фамилия</th>
                                <th class="text-center">Имя</th>
                                <th class="text-center">Отчество</th>
                                <th class="text-center"></th>
                                <th class="text-center"></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="item in students_l" :key="item.name">
                                <td class="text-center">{{ item.id }}</td>
                                <td class="text-center">{{ item.surname }}</td>
                                <td class="text-center">{{ item.middleName }}</td>
                                <td class="text-center">{{ item.name }}</td>
                                <td class="text-center">
                                    <v-col>
                                    <v-text-field style="max-width: 50px"
                                        label="№ ПК"
                                        v-model="item.comp"
                                    ></v-text-field>
                                    </v-col>
                                </td>
                                <td class="text-center">
                                    <v-btn small
                                    @click="Use(item)"
                                    style="margin-right: 10px " color="success">Подключить</v-btn>
                                    <v-btn small color="error"
                                    @click="Quit(item)"
                                    >Отключить</v-btn>
                                </td>
                            </tr>
                        </tbody>
                    </template>
                </v-simple-table>
            </v-card-text>
        </v-card>
        <v-dialog v-model="dialog_templ" persistent max-width="600px">
            <v-card>
                <v-card-title>
                <span class="headline">Задайте имя шаблону</span>
                </v-card-title>
                <v-card-text>
                <v-container>
                    <v-text-field v-model="temple.name" label="Название" required></v-text-field>
                </v-container>
                </v-card-text>
                <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="red darken-1" text @click="dialog_templ = false">Отмена</v-btn>
                <v-btn color="blue darken-1" text @click="SaveTamplOk">Сохранить</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-container>
</template>

<script>
export default {
    name: "TableComps",
    data() {
        return {
            ROOM_LIST: [
                "1",
                "2",
                "3"
            ],
            temple: {
                name: "",
            },
            dialog_templ: false,
            students_l: []
        }
    },
    watch: {
        students: function () {
            this.students_l = this.students;
        }
    },
    computed: {
        group: {
            get() {
                return this.$store.state.students.group
            },
            set(v) {
                this.$store.commit("GROUP_COMP", v)
            }
        },
        room: {
            get() {
                return this.$store.state.students.room
            },
            set(v) {
                this.$store.commit("ROOM_COMP", v)
            }
        },
        students: {
            get() {
                return this.$store.state.students.students
            }
        },
        GROUP_LIST() {
            return this.$store.state.students.GROUP_LIST
        },
        STUDENTS_SELECT_GROUP() {
            return this.$store.state.students.STUDENTS_SELECT_GROUP
        },
    },
    created() {
        this.$store.dispatch("GET_GROUPS");
        this.$store.dispatch("GET_STUDENTS_SELECT_GROUP", this.group)
        .then(() => {
            let studs = [];
            this.STUDENTS_SELECT_GROUP.map((i) => {
                let st = {
                    id: i.id,
                    surname: i.surname,
                    name: i.name,
                    middleName: i.middleName,
                    numGroup: i.numGroup,
                    login: i.login,
                    comp: ""
                }
                studs.push(st);
            })
            this.$store.commit("STUDENTS_COMP", studs)
            this.students_l = this.students;
        })
    },
    methods: {
        selectGroup() {
            this.$store.dispatch("GET_STUDENTS_SELECT_GROUP", this.group)
            .then(() => {
            let studs = [];
            this.STUDENTS_SELECT_GROUP.map((i) => {
                let st = {
                    id: i.id,
                    surname: i.surname,
                    name: i.name,
                    middleName: i.middleName,
                    numGroup: i.numGroup,
                    login: i.login,
                    comp: ""
                }
                studs.push(st);
            })
            this.$store.commit("STUDENTS_COMP", studs)
            this.students_l = this.students;
        })
        },
        Quit(item) {
            let obj = {
                login: item.login,
                numGroup: this.group,
                numRoom: this.room,
                numComp: item.comp
            }
            this.$store.dispatch("TurnOFFPC", obj)
        },
        Use(item) {
            let obj = {
                login: item.login,
                numGroup: this.group,
                numRoom: this.room,
                numComp: item.comp
            }
            this.$store.dispatch("TurnONPC", obj)
        },
        UseAll() {
            this.students.map((s) => {
                let obj = {
                    login: s.login,
                    numGroup: this.group,
                    numRoom: this.room,
                    numComp: s.comp
                }
                this.$store.dispatch("TurnONPC", obj)
            })
        },
        QuitAll() {
            this.students.map((s) => {
                let obj = {
                    login: s.login,
                    numGroup: this.group,
                    numRoom: this.room,
                    numComp: s.comp
                }
                this.$store.dispatch("TurnOFFPC", obj)
            })
        },
        SaveTempl() {
            this.dialog_templ = true
        },
        SaveTamplOk() {
            let templ = {
                id: 0,
                name: this.temple.name,
                room: this.room,
                group: this.group,
                students: this.students_l
            }
            this.$store.commit("UPDATE_TEMPLATES_COMP", templ)
            this.temple = {
                name: "",
            },
            this.dialog_templ = false
        }
    }
}
</script>

<style scoped>
td {
    padding: 0 !important;
}
.col {
    padding: 0 !important;
}
</style>