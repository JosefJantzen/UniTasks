import {createStore} from 'vuex'

const store = createStore({
    state () {
        return {
        test: 0
        }
    },
    mutations: {
        increment (state) {
        state.test++
        }
    }
})

export default store