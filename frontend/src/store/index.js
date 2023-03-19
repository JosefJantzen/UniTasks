import {createStore} from 'vuex'
import VuexPersistence from 'vuex-persist'

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
                    if (!recTask.history[i].done) {
                        recTask.history[i].count = parseInt(i)
                        recTask.history[i].countMax = recTask.history.length
                        recTask.history[i].name = recTask.name
                        recTask.history[i].recurring = true
                        res.push(recTask.history[i])
                        break
                    }
                }
            }      
            return [...tasks, ...res]
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