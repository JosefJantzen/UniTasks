import axios from 'axios'

const baseURL = process.env.VUE_APP_API_ENDPOINT

axios.defaults.withCredentials = true

const restClient = () => {

    return axios.create({
        baseURL: baseURL,
        
    })
}

export default {
    get: async (path, params) => {
        let cli = await restClient()
        return cli.get(path, { params: params })
    },
    post: async (path, data) => {
        let cli = await restClient()
        return cli.post(path, data)
    },
    put: async (path, data) => {
        let cli = await restClient()
        return cli.put(path, data)
    },
    delete: async (path) => {
        let cli = await restClient()
        return cli.delete(path)
    }
}