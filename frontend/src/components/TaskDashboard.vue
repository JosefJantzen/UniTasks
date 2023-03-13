<template>
    <va-tabs v-model="currentTab" center>
        <template #tabs>
            <va-tab
                v-for="tab in ['List view', 'Calendar view']"
                :key="tab"
                class="tab"
            >
                {{ tab }}
            </va-tab>
        </template>
    </va-tabs>
    <br>
    <TaskList v-if="currentTab == 0"></TaskList>
    <TaskCalendar v-else-if="currentTab == 1"></TaskCalendar>
</template>

<script>
import { mapActions } from 'vuex'

import TaskList from './dashboard/TaskList.vue'
import TaskCalendar from './dashboard/TaskCalendar.vue'

export default {
    name: 'TaskDashboard',
    components: {
        TaskList,
        TaskCalendar
    },
    methods: {
        ...mapActions('tasks', ['list']),
        foo () {
            this.list()
        }
    },
    data: () => ({
        currentTab: 0,
    })
}

</script>

<style>
.tab {
    padding-left: 1rem;
    padding-right: 1rem;
}
</style>