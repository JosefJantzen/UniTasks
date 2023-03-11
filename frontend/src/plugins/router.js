import {createRouter, createWebHashHistory} from 'vue-router'

import TaskDashboard from '../components/TaskDashboard.vue'
import Login from '../components/Login.vue'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        { path: '/', component: TaskDashboard},
        { path: '/login', component: Login}
    ],
})

export default router