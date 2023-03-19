import { createApp } from 'vue'

import store from './store/index'
import router from './plugins/router'
import vuestic from './plugins/vuestic'

import App from './App.vue'

createApp(App)
	.use(vuestic)
	.use(router)
	.use(store.store)
	.mount('#app')

setInterval(() => {
	if (getCookie("token") == "") {
		//store.clear()
	}
}, 500);

function getCookie(cname) {
	let name = cname + "=";
	let decodedCookie = decodeURIComponent(document.cookie);
	let ca = decodedCookie.split(';');
	for(let i = 0; i <ca.length; i++) {
		let c = ca[i];
		while (c.charAt(0) == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) == 0) {
			return c.substring(name.length, c.length);
		}
	}
	return "";
  }