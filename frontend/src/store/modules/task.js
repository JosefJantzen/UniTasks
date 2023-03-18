//import moment from 'moment'
import api from '../../api/apiClient'

const state = () => ({
    tasks: []
})

const getters = {
    getAll: (state) => {
        return state.tasks //.sort((a, b) => moment(String(a.due)) - moment(String(b.due)))
    },
    getById: (state, id) => {
        for (const task of state.tasks) {
            if (id == task.id) {
                return task
            }
        }
        return null
    }
}

const mutations = {
    clear: (state) => {
        state.tasks = []
    },
    add(state, task) {
        state.tasks.push(task)
    },
}

const actions = {
    list: async (context) => {
        await api.get('/tasks').then((res) => {
            context.commit('clear')
            for (const task of res.data) {
                task.recurring = false
                context.commit('add', task)
            }
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