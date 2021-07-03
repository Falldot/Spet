import {HTTP} from '../aixos/http'

const
rus = "щ   ш  ч  ц  ю  я  ё  ж  ъ  ы  э  а б в г д е з и й к л м н о п р с т у ф х ь".split(/ +/g),
eng = "shh sh ch cz yu ya yo zh `` y' e` a b v g d e z i j k l m n o p r s t u f x `".split(/ +/g);

function gen_password(){
    var password = "";
    var symbols = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    for (var i = 0; i < 11; i++){
        password += symbols.charAt(Math.floor(Math.random() * symbols.length));
    }
    return password;
}

function translit(text, engToRus) {
    var x;
    try {
        for(x = 0; x < rus.length; x++) {
            text = text.split(engToRus ? eng[x] : rus[x]).join(engToRus ? rus[x] : eng[x]);
            text = text.split(engToRus ? eng[x].toUpperCase() : rus[x].toUpperCase()).join(engToRus ? rus[x].toUpperCase() : eng[x].toUpperCase());
        }
    } catch (error) {
        console.log(error)
        return "";
    }
    return text;
}

export default {
    state: {
        group: JSON.parse(localStorage.getItem("group")) || "",
        room: JSON.parse(localStorage.getItem("room")) || "",
        students: [],
        TEMP: [
            {}
        ],
        TEMPLATES: JSON.parse(localStorage.getItem("Templates")) || [],
        TEMPLATES_COMPS: JSON.parse(localStorage.getItem("Templates_Comps")) || [],
        HEADERS: {
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
            info: false,
            activs: false,
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
        FILTRES:
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
            activs: "",
            phone: "",
            info: "",
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
        },
        ITEMS: {
            dop: ["Сирота", "Инвалид", "Мал. Обес.", "ИДН", "КДН", "Многодет.", "Непол. семья"],
            budget: ["Бюджетная", "Платная"],
            gender: ["М", "Ж"],
            status: ["Учиться", "Не учиться"]
        },
        STUDENTS: [
        ],
        SORT_STUDENTS: [
        ],
        GROUP_LIST: [
        ],
        LOADING: true,
        SELECTED: [],
        TABLE_HEADERS: [
            { text: 'Фамилия', value: 'surname' },
            { text: 'Имя', value: 'name' },
            { text: 'Отчество', value: 'middleName' },
            { text: 'Группа', value: 'numGroup' },
        ],
        STUDENTS_SELECT_GROUP: []
    },
    mutations: {
        STUDENTS_COMP(state, items) {
            state.students = items
        },
        ROOM_COMP(state, items) {
            let sItem = JSON.stringify(items);
            localStorage.setItem("room", sItem);
            state.room = items
        },
        GROUP_COMP(state, items) {
            let sItem = JSON.stringify(items);
            localStorage.setItem("group", sItem);
            state.group = items
        },
        STUDENTS_COMP_TEMPL(state, items) {
            state.students = items.students
            state.room = items.room
            state.group = items.group
        },
        UPDATE_STUDENTS(state, students) {
            state.STUDENTS = students
        },
        UPDATE_STUDENTS_SELECT_GROUP(state, students) {
            state.STUDENTS_SELECT_GROUP = students
        },
        UPDATE_FILTRES(state, filtres) {
            state.FILTRES = filtres
        },
        LOADING(state, bool) {
            state.LOADING = bool
        },
        UPDATE_TEMPLATES(state, item) {
            item.id = state.TEMPLATES.length;
            state.TEMPLATES.push(item);
            let sItem = JSON.stringify(state.TEMPLATES);
            localStorage.setItem("Templates", sItem);
            state.TEMPLATES = JSON.parse(localStorage.getItem("Templates"));
        },
        DELETE_TEMPLATES(state, id) {
            state.TEMPLATES = state.TEMPLATES.filter(i => {
                return i.id !== id
            })
            let sItem = JSON.stringify(state.TEMPLATES);
            localStorage.setItem("Templates", sItem);
        },
        UPDATE_TEMPLATES_COMP(state, item) {
            item.id = state.TEMPLATES_COMPS.length;
            state.TEMPLATES_COMPS.push(item);
            let sItem = JSON.stringify(state.TEMPLATES_COMPS);
            localStorage.setItem("Templates_Comps", sItem);
            state.TEMPLATES_COMPS = JSON.parse(localStorage.getItem("Templates_Comps"));
        },
        DELETE_TEMPLATES_COMP(state, id) {
            state.TEMPLATES_COMPS = state.TEMPLATES_COMPS.filter(i => {
                return i.id !== id
            })
            let sItem = JSON.stringify(state.TEMPLATES_COMPS);
            localStorage.setItem("Templates_Comps", sItem);
        },
        UPDATE_HEADERS(state, headers) {
            state.HEADERS = headers
        },
        UPDATE_SORT_STUDENTS(state, students) {
            state.SORT_STUDENTS = students
        },
        UPDATE_SELECTED(state, students) {
            state.SELECTED = students
        },
        UPDATE_HEADERS_TABLES(state, {headers, filtres}) {
            state.TABLE_HEADERS = []
            for (let i in headers) {
                if (i === "id" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'ID', value: 'id' }
                )
                if (i === "surname" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Фамилия', value: 'surname' }
                )
                if (i === "name" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Имя', value: 'name' }
                )
                if (i === "middleName" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Отчетсво', value: 'middleName' }
                )
                if (i === "date_b" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Дата рождения', value: 'date_b' }
                )
                if (i === "city" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Нас. пункт', value: 'city' }
                )
                if (i === "street" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Улица', value: 'street' }
                )
                if (i === "house" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Дом', value: 'house' }
                )
                if (i === "flat" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Квартира', value: 'flat' }
                )
                if (i === "phone" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Дом. телефон', value: 'phone' }
                )
                if (i === "info" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Доп. информация', value: 'info' }
                )
                if (i === "gender" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Пол', value: 'gender' }
                )
                if (i === "numGroup" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Группа', value: 'numGroup' }
                )
                if (i === "status" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Статус', value: 'status' }
                )
                if (i === "login" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Логин', value: 'login' }
                )
                if (i === "password" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Пароль', value: 'password' }
                )
                if (i === "budget" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Основа', value: 'budget' }
                )
                if (i === "orphan" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Сирота', value: 'orphan' }
                )
                if (i === "invalid" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Инвалид', value: 'invalid' }
                )
                if (i === "low_income" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Мал. обес.', value: 'low_income' }
                )
                if (i === "low_parents" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Непол. семья', value: 'low_parents' }
                )
                if (i === "idn" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'ИДН', value: 'idn' }
                )
                if (i === "kdn" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'КДН', value: 'kdn' }
                )
                if (i === "many_children" & headers[i]) state.TABLE_HEADERS.push(
                    { text: 'Многодет.', value: 'many_children' }
                )
            }
            let ArrayMain = [];

            for (let el in state.STUDENTS) {
                let add = true;
                let addfull = false;
                for (let i in state.STUDENTS[el]) {
                    if (filtres[i] != "")
                    {
                        if (Array.isArray(filtres[i])) {
                            for(let arrItem in filtres[i]) {
                                if (state.STUDENTS[el][i] != filtres[i][arrItem]) {
                                    add = false;
                                } else {
                                    addfull = true;
                                }
                            }
                        }
                        if (state.STUDENTS[el][i] != filtres[i]) {
                            add = false;
                        }
                    }
                }
                if (add) {
                    ArrayMain.push(state.STUDENTS[el]);
                } else if (addfull) {
                    ArrayMain.push(state.STUDENTS[el]);
                }
            }
            state.SORT_STUDENTS = ArrayMain
        }
    },
    actions: {
        UPDATE_HEADERS_TABLES_LITE({commit, state}) {
            commit("UPDATE_FILTRES", {
                filtres: state.FILTRES
            })
            commit("UPDATE_HEADERS", {
                headers: state.HEADERS
            })
        },
        async TurnONPC({commit}, item) {
            commit('LOADING', true)
            item.numGroup = translit(item.numGroup);
            await HTTP.post('turnONPC', item)
            .then(() => {
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async TurnOFFPC({commit}, item) {
            commit('LOADING', true)
            item.numGroup = translit(item.numGroup);
            await HTTP.post('turnOFFPC', item)
            .then(() => {
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async ADD_STUDENT({commit, state, dispatch}, item) {
            commit('LOADING', true)
            item.login = (translit(item.surname) + "_" + translit(item.name[0]) + translit(item.middleName[0])).toLowerCase();
            item.password = gen_password();
            await HTTP.post('students', item)
            .then(({data}) => {
                state.STUDENTS.push(data)
                dispatch('UPDATE_HEADERS_TABLES_LITE')
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async ADD_STUDENT_ALL({commit, dispatch}, data) {
            commit('LOADING', true)
            data.forEach((item) => {
                if (item[0] == "Фамилия" | item[0] == "") {
                    console.log("")
                } else {
                    let obj = {
                        id: "",
                        surname: item[0],
                        name: item[1],
                        middleName: item[2],
                        date_b: "",
                        city: "",
                        street: "",
                        house: "",
                        flat: "",
                        activs: "",
                        phone: "",
                        info: "",
                        gender: "",
                        numGroup: item[3],
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
                        many_children: false
                    }
                    dispatch("ADD_STUDENT", obj)
                }
            });
            commit('LOADING', false)
        },
        async GET_STUDENTS({commit}) {
            commit('LOADING', true)
            await HTTP.get('students')
            .then(({data}) => {
                commit('UPDATE_STUDENTS', data)
                commit('UPDATE_SORT_STUDENTS', data)
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async GET_STUDENTS_SELECT_GROUP({commit}, name) {
            commit('LOADING', true)
            await HTTP.post('studentsGroup', {name})
            .then(({data}) => {
                commit('UPDATE_STUDENTS_SELECT_GROUP', data)
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async UPDATE_STUDENT({commit, state, dispatch}, item) {
            commit('LOADING', true)
            await HTTP.put('students', item)
            .then(({data}) => {
                state.STUDENTS = state.STUDENTS.map(o => {
                    if (o.id === data.id) {
                        return data;
                    }
                    return o;
                });
                dispatch('UPDATE_HEADERS_TABLES_LITE')
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async DELETE_STUDENT({commit, state, dispatch}) {
            commit('LOADING', true)
            for (let v in state.SELECTED) {
                await HTTP.post('studentsDel', state.SELECTED[v])
                .then(() => {
                    state.SELECTED.map(a => {
                        state.STUDENTS = state.STUDENTS.filter(i => {
                            return i.id !== a.id
                        })
                        state.SORT_STUDENTS = state.SORT_STUDENTS.filter(i => {
                            return i.id !== a.id
                        })
                    })
                    dispatch('UPDATE_HEADERS_TABLES_LITE')
                    commit('LOADING', false)
                })
                .catch((err) => {
                    console.log(err)
                    commit('LOADING', false)
                })
            }
            state.SELECTED = []
        },
        async CREATE_GROUP({commit, state}, name) {
            commit('LOADING', true)
            await HTTP.post('groups', {name})
            .then(() => {
                state.GROUP_LIST.push(name)
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async GET_GROUPS({commit, state}) {
            state.GROUP_LIST = [];
            commit('LOADING', true)
            await HTTP.get('groups')
            .then(({data}) => {
                for (let v in data) {
                    state.GROUP_LIST.push(data[v].name)
                }
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
        async DELETE_GROUP({dispatch, commit}, name) {
            commit('LOADING', true)
            await HTTP.post('groupsDel', {name})
            .then(() => {
                dispatch("GET_GROUPS");
                commit('LOADING', false)
            })
            .catch((err) => {
                console.log(err)
                commit('LOADING', false)
            })
        },
    }
}