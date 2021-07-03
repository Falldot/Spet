//import {HTTP} from '../aixos/http'

export default {
    state: {
        from: {},
        login: localStorage.getItem("login") || false,
        loginName: localStorage.getItem("loginName") || "",
        password: localStorage.getItem("password") || "",
        role: localStorage.getItem("role") || 3,
        save: false,
    },
    mutations: {
        UPDATE_USER(state, user) {
            state.loginName = user.loginName;
            state.password = user.password;
            state.role = user.role;
            state.login = true
        },
        LOGOUT(state) {
            state.loginName = "";
            state.password = "";
            state.role = 0;
            state.login = false
        },
        SET_SAVE(state, bool) {
            state.save = bool;
        }
    },
    actions: {
        async LOGIN({commit}, input) {
            if (input.loginName == "admin" & input.password == "admin") {
                commit("SET_SAVE", true);
                commit('UPDATE_USER', {
                    loginName: input.loginName,
                    password: input.password,
                    role: 1,
                })
                if (input.save) {
                    localStorage.setItem("login", true);
                    localStorage.setItem("loginName", input.loginName);
                    localStorage.setItem("password", input.password);
                    localStorage.setItem("role", 1);
                }
            }
        },
        async LOGOUT({commit}) {
            commit('LOGOUT');
        }
    }
}