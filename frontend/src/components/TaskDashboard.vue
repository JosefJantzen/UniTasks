<template>
    <va-tabs v-model="currentTab" center style="position: relative;">
        <template #tabs>
            <va-tab
                v-for="tab in ['Pending Tasks', 'One time tasks', 'Recurring tasks']"
                :key="tab"
                class="tab"
            >
                {{ tab }}
            </va-tab>
        </template>
        <va-switch
            v-model="showDone"
            left-label
            off-color="info"
            color="secondary"
            true-inner-label="Show done"
            false-inner-label="Hide done"
            style="position: absolute; top: 0; right: 5%; float: left;"
        />
    </va-tabs>
    <div class="view">
        <PendingTasks v-if="currentTab == 0" :showDone="this.showDone"></PendingTasks>
        <OneTimeTasks v-else-if="currentTab == 1" :showDone="this.showDone"></OneTimeTasks> 
        <RecurringTasks v-else-if="currentTab == 2" :showDone="this.showDone"></RecurringTasks>
    </div>
    
</template>

<script>
import PendingTasks from './dashboard/PendingTasks.vue'
import OneTimeTasks from './dashboard/OneTimeTasks.vue'
import RecurringTasks from './dashboard/RecurringTasks.vue'

export default {
    name: 'TaskDashboard',
    components: {
        PendingTasks,
        OneTimeTasks,
        RecurringTasks
    },
    methods: {
        
    },
    data: () => ({
        currentTab: 0,
        showDone: false,
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