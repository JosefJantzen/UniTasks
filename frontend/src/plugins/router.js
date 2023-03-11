import {createRouter, createWebHashHistory} from 'vue-router'
import api from '../api/apiClient'

import TaskDashboard from '../components/TaskDashboard.vue'
import Login from '../components/Login.vue'

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
        { path: '/', component: TaskDashboard, beforeEnter: requireAuth},
        { path: '/login', component: Login}
    ],
})

export default router