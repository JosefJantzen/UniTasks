import api from '../../api/apiClient'

const state = () => ({
    user: null
})

const getters = {
    get: (state) => {
        return state.user
    }
}

const mutations = {
    clear: (state) => {
        state.user = null
    },
    set: (state, user) => {
        state.user = user
    },
    setProperty: (state, data) => {
        state[data.property] = data.value
    }
}

const actions = {
    signIn: async (context, credentials)  => {
        await api.post('/signIn', credentials).then(() => {
            delete credentials.pwd
            context.commit('set', credentials)
        }).catch((e) => {
            throw e
        })          
    },
    signUp: async (context, credentials) => {
        await api.post('/signUp', credentials).then(() => {
            delete credentials.pwd
            context.commit('set', credentials)
        }).catch((e) => {
            throw e
        }) 
    }
}

export default {
    namespaced: true,
    
    state: state,
    getters: getters,
    mutations: mutations,
    actions: actions
    
}