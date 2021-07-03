import axios from 'axios'

export const HTTP = axios.create({
    //baseURL: "http://localhost:3000/"
    baseURL: "http://local.spet.pro:15280/"
})