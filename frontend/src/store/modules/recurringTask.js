import api from '../../api/apiClient'

const state = () => ({
    recurringTasks: []
})

const getters = {
    getAll: (state) => {
        return state.recurringTasks.sort((a, b) => moment(String(a.ending)) - moment(String(b.ending)))
    },
    getById: (state, id) => {
        for (const task of state.recurringTasks) {
            if (id == task.id) {
                return task
            }
        }
        return null
    }
}

const mutations = {
    clear: (state) => {
        state.recurringTasks = []
    },
    add(state, task) {
        state.recurringTasks.push(task)
    },
}

const actions = {
    listRecurring: async (context) => {
        await api.get('/recurring-tasks').then((res) => {
            context.commit('clear')
            for (const task of res.data) {
                task.recurring = true
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