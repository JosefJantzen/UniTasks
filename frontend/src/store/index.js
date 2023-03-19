import {createStore} from 'vuex'
import VuexPersistence from 'vuex-persist'
//import moment from 'moment'

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
            //let recTasks = store.getters['recurringTasks/getAll'] 
            /*for (const recTask in recTasks) {
                recTask.history.sort((a, b) => moment(String(a.due)) - moment(String(b.due)))
                console.log(recTask)
            }*/        
            return tasks
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