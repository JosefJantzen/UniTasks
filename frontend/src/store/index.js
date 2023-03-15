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
        getAllTasks: () => {
            let tasks = store.getters['tasks/getAll']
            let recTasks = store.getters['recurringTasks/getAll']            
            return [ ...recTasks, ...tasks]
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