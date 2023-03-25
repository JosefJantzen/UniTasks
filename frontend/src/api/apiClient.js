import axios from 'axios'

const baseURL = process.env.VUE_APP_API_ENDPOINT

axios.defaults.withCredentials = true

export default {
    get: async (path) => {
        return axios.get(baseURL + path)
    },
    post: async (path, data) => {
        return axios.post(baseURL + path, data)
    },
    put: async (path, data) => {
        return axios.put(baseURL + path, data)
    },
    delete: async (path, payload) => {
        axios.delete(baseURL + path, { data: payload})
    }
}