import {createRouter, createWebHashHistory} from 'vue-router'
import api from '../api/apiClient'

import Login from '../components/Login.vue'
import SignUp from '../components/SignUp.vue'
import TaskDashboard from '../components/TaskDashboard.vue'
import Settings from '../components/Settings.vue'
import NewTask from '../components/NewTask.vue'

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
  }

function requireAuth (to, from, next) {
    const token = getCookie("token")
    if (token) {
        api.get('/refresh')
        next()
        return
    }
    next({
        path: '/login',
        query: { redirect: to.fullPath }
    })
}

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        { path: '/login', component: Login},
        { path: '/signUp', component: SignUp},

        { path: '/', component: TaskDashboard, beforeEnter: requireAuth},
        { path: '/settings', component: Settings, beforeEnter: requireAuth},
        { path: '/newTask', component: NewTask, beforeEnter: requireAuth}
    ],
})

export default router