import {createStore} from 'vuex'
import VuexPersistence from 'vuex-persist'

import user from './modules/user'

const vuexLocal = new VuexPersistence({
    storage: window.localStorage
  })

const store = createStore({
    modules: {
        user: user
    },
    plugins: [vuexLocal.plugin]
})


export default {
    store,
    clear () {
        store.commit('user/clear')
    }
}