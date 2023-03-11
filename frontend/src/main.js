import { createApp } from 'vue'

import store from './store/index'
import router from './plugins/router'
import vuestic from './plugins/vuestic'

import App from './App.vue'

createApp(App)
	.use(vuestic)
	.use(router)
	.use(store)
	.mount('#app')
