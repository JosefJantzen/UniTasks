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
        Vue.set(state, 'user', null)
    },
    set: (state, user) => {
        state.user = user
    },
    setProperty: (state, data) => {
        Vue.set(state.user, data.property, data.value)
    }
}

const actions = {
    signIn: async (context, credentials)  => {
        try {
            let res = await api.post('/signIn', credentials)
            delete credentials.pwd
            context.commit('set', credentials)
        }
        catch (e) {
            throw e
        }
    }
}

export default {
    namespace: true,
    state: state,
    getters: getters,
    mutations: mutations,
    actions: actions
}