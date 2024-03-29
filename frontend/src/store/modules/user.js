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
        state.user[data.property] = data.value
    }
}

const actions = {
    async signIn (context, credentials)  {
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
    },
    logout: (context) => {
        clearInterval(context.state.intervalId)
        document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/; Secure; SameSite=None; Domain=" + window.location.hostname + ";";
        location.reload()
    },
    changeMail: async (context, mail) => {
        await api.put('/updateMail', {
            eMail: mail
        }).then(() => {
            context.commit('setProperty', {
                property: 'eMail',
                value: mail
            })
        }).catch((e) => {
            throw e
        })
    },
    changePwd: async (context, pwd) => {
        await api.put('/updatePwd', {
            pwd: pwd
        }).then(() => {

        }).catch((e) => {
            throw e
        })
    },
    async refresh  ()  {
        await api.get('/refresh')
    },
    interval: (context, id) => {
        context.commit('setProperty', {
            property: 'intervalId',
            value: id
        })
    },
    deleteUser: async (context, credentials) => {
        await api.delete('/deleteUser', credentials).then(() => {
            context.dispatch('logout')
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