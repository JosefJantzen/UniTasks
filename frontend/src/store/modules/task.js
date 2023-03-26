import moment from 'moment'
import api from '../../api/apiClient'

const state = () => ({
    tasks: []
})

const getters = {
    getAll: (state) => {
        return state.tasks.sort((a, b) => {
            if (a.done && !b.done) {
                return 1
            }
            else if (!a.done && b.done) {
                return -1
            }
            return moment(String(a.due)) - moment(String(b.due))
        })
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
    add: (state, task) => {
        state.tasks.push(task)
    },
    update: (state, task) => {
        const i = state.tasks.findIndex(t => t.id == task.id)
        state.tasks[i] = task
    },
    delete: (state, id) => {
        const i = state.tasks.findIndex(t => t.id == id)
        state.tasks.splice(i, 1)
    }
}

const actions = {
    listTask: async (context) => {
        await api.get('/tasks').then((res) => {
            context.commit('clear')
            for (const task of res.data) {
                task.recurring = false
                context.commit('add', task)
            }
        }).catch((e) => {
            throw e
        })
    },
    done: async (context, task) => {
        await api.put('/tasks/' + task.id + '/done', {
            "done": task.done,
            "doneAt": task.doneAt
        }).then(() => {
            context.commit('update', task)
        }).catch((e) => {
            throw e
        })
    },
    createTask: async (context, task) => {
        await api.post('/tasks', task).then((res) => {
            task.id = res.data.id
            context.commit('add', task)
        }).catch((e) => {
            throw e
        })
    },
    updateTask: async (context, task) => {
        await api.put('/tasks/' + task.id, task).then(() => {
            context.commit('update', task)
        }).catch((e) => {
            throw e
        })
    },
    deleteTask: async (context, task) => {
        await api.delete('/tasks/' + task.id).then(() => {
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