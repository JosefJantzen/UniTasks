import moment from 'moment'
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
    add: (state, task) => {
        state.recurringTasks.push(task)
    },
    update: (state, task) => {
        const i = state.recurringTasks.findIndex(t => t.id == task.id)
        state.recurringTasks[i] = task
    },
    updateHist: (state, hist) => {
        const i = state.recurringTasks.findIndex(t => t.id == hist.recurringTaskId)
        const ii = state.recurringTasks[i].history.findIndex(t => t.id == hist.id)
        state.recurringTasks[i].history[ii] = hist
    },
    delete: (state, id) => {
        const i = state.recurringTasks.findIndex(t => t.id == id)
        state.recurringTasks.splice(i, 1)
    }
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
    },
    createRecurring: async (context, task) => {
        await api.post('/recurring-tasks', task).then((res) => {
            task.id = res.data.id
            context.commit('add', task)
        }).catch((e) => {
            throw e
        })
    },
    updateRecurring: async (context, task) => {
        await api.put('/recurring-tasks/' + task.id, task).then(() => {
            context.commit('update', task)
        }).catch((e) => {
            throw e
        })
    },
    doneHist: async (context, task) => {
        await api.put('/recurring-tasks-history/' + task.id + '/done', {
            done: task.done,
            recurringTaskId: task.recurringTaskId
        }).then(() => {
            context.commit('updateHist', task)
        }).catch((e) => {
            throw e
        })
    },
    deleteRecurring: async (context, task) => {
        await api.delete('/recurring-tasks/' + task.id).then(() => {
            context.commit('delete', task.id)
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