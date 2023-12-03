import axios from 'axios'

// const baseURL = "http://127.0.0.1:8080/api";
const baseURL = "/api";

const api = axios.create({
    baseURL: baseURL,
    withCredentials: false,
})


export default api;
