import {createStore} from 'vuex'
import VuexPersistence from 'vuex-persist'
import moment from 'moment'

import user from './modules/user'
import tasks from './modules/task'
import recurringTasks from './modules/recurringTask'

const vuexLocal = new VuexPersistence({
    storage: window.localStorage
  })

const store = createStore({
    modules: {
        user: user,
        tasks: tasks,
        recurringTasks: recurringTasks
    },
    getters: {
        getPendingTasks: () => {
            let tasks = store.getters['tasks/getAll']
            let recTasks = store.getters['recurringTasks/getAll']
            let res = [] 
            for (const recTask of recTasks) {
                for (const i in recTask.history) {
                        recTask.history[i].count = parseInt(i) + 1
                        recTask.history[i].countMax = recTask.history.length
                        recTask.history[i].name = recTask.name
                        recTask.history[i].recurring = true
                }
            } 
            for (const recTask of recTasks) {
                for (const i in recTask.history) {
                    res.push(recTask.history[i])
                    if (!recTask.history[i].done) {
                        break
                    }
                }
            } 
            return [...tasks, ...res].sort((a, b) => {
                if (a.done && !b.done) {
                    return 1
                }
                else if (!a.done && b.done) {
                    return -1
                }
                return moment(String(a.due)) - moment(String(b.due))
            })
        }
    },
    plugins: [vuexLocal.plugin]
})


export default {
    store,
    clear () {
        store.commit('user/clear')
        store.commit('tasks/clear')
        store.commit('recurringTasks/clear')
    }
}