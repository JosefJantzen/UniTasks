<template>
    <va-tabs v-model="currentTab" center>
        <template #tabs>
            <va-tab
                v-for="tab in ['Pending Tasks', 'One time tasks', 'Recurring tasks']"
                :key="tab"
                class="tab"
            >
                {{ tab }}
            </va-tab>
        </template>
    </va-tabs>
    <br>
    <div class="view">
       <PendingTasks v-if="currentTab == 0"></PendingTasks>
        <OneTimeTasks v-else-if="currentTab == 1"></OneTimeTasks> 
    </div>
    
</template>

<script>
import PendingTasks from './dashboard/PendingTasks.vue'
import OneTimeTasks from './dashboard/OneTimeTasks.vue'

export default {
    name: 'TaskDashboard',
    components: {
        PendingTasks,
        OneTimeTasks
    },
    methods: {
        
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

.view {
    padding-top: 3%;
    margin-left: 15%;
    margin-right: 15%;
}
</style>