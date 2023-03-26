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
		store.clear()
	}
}, 500);

window.onload = async function () {
	try {
		await store.store.dispatch('user/refresh')
		let intervalId = setInterval(async () =>{
			try {
				await store.store.dispatch('user/refresh')
			}
			catch (e) {
				if(e.response.status == 401) {
					router.push('/login')
					clearInterval(intervalId)
				}
			}
		}, 270000)
		store.store.dispatch('user/interval', intervalId)
	}
	catch (e) {
		if(e.response.status == 401) {
			router.push('/login')
		}
		return
	}
}

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