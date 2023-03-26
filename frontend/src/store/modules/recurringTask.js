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
    addHist: (state, hist) => {
        const i = state.recurringTasks.findIndex(t => t.id == hist.recurringTaskId)
        state.recurringTasks[i].history.push(hist)
        state.recurringTasks[i].history.sort((a, b) => moment(String(a.due)) - moment(String(b.due)))
    },
    updateHist: (state, hist) => {
        const i = state.recurringTasks.findIndex(t => t.id == hist.recurringTaskId)
        const ii = state.recurringTasks[i].history.findIndex(t => t.id == hist.id)
        state.recurringTasks[i].history[ii] = hist
    },
    delete: (state, id) => {
        const i = state.recurringTasks.findIndex(t => t.id == id)
        state.recurringTasks.splice(i, 1)
    },
    deleteHist: (state, hist) => {
        const i = state.recurringTasks.findIndex(t => t.id == hist.recurringTaskId)
        const ii = state.recurringTasks[i].history.findIndex(t => t.id == hist.id)
        state.recurringTasks[i].history.splice(ii, 1)
    }
}

const actions = {
    listRecurring: async (context) => {
        await api.get('/recurring-tasks').then((res) => {
            context.commit('clear')
            for (const task of res.data) {
                task.recurring = true
                task.history.sort((a, b) => moment(String(a.due)) - moment(String(b.due)))
                context.commit('add', task)
            }
        }).catch((e) => {
            throw e
        })
    },
    createRecurring: async (context, task) => {
        return await api.post('/recurring-tasks', task).then((res) => {
            task.id = res.data.id
            context.commit('add', task)
            return res.data.id
        }).catch((e) => {
            throw e
        })
    },
    createRecurringHist: async (context, task) => {
        await api.post('/recurring-tasks-history', task).then((res) => {
            task.id = res.data.id
            context.commit('addHist', task)
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
    updateRecurringHist: async (context, task) => {
        await api.put('/recurring-tasks-history/' + task.id, task).then(() => {
            context.commit('updateHist', task)
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
    },
    deleteRecurringHist: async (context, task) => {
        await api.delete('/recurring-tasks-history/'  + task.id).then(() => {
            context.commit('deleteHist', task)
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