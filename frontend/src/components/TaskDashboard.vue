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
    <div class="view">
       <TaskList v-if="currentTab == 0"></TaskList>
        <TaskCalendar v-else-if="currentTab == 1"></TaskCalendar> 
    </div>
    
</template>

<script>
import TaskList from './dashboard/TaskList.vue'
import TaskCalendar from './dashboard/TaskCalendar.vue'

export default {
    name: 'TaskDashboard',
    components: {
        TaskList,
        TaskCalendar
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