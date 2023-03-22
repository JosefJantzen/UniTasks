<template>
    <div class="main">
        <va-tabs v-model="currentTab" center style="position: fixed; width: 100%;">
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
                :class="this.isMobile ? 'switch-mobile' : 'switch'"
            />
        </va-tabs>
    </div>
    
    <div class="view" :style="this.isMobile ? 'margin-top: 150px;' : 'margin-top: 105px;'">
        <PendingTasks v-if="currentTab == 0" :showDone="this.showDone"></PendingTasks>
        <OneTimeTasks v-else-if="currentTab == 1" :showDone="this.showDone"></OneTimeTasks> 
        <RecurringTasks v-else-if="currentTab == 2" :showDone="this.showDone"></RecurringTasks>
    </div>
    
</template>

<script>
import help from '../help/help'

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
        isMobile: help.isMobile()
    })
}

</script>

<style>
.main {
    position: relative;
    top: 100px;
    width: 100;
    overflow: hidden;
    z-index: 200;
}

.tab {
    padding-left: 1rem;
    padding-right: 1rem;
}

.view {
    padding-top: 3%;
    margin-left: 15%;
    margin-right: 15%;
}

.switch {
    position: fixed;
    top: 100px;
    right: 5%;
    float: left;
}

.switch-mobile {
    position: fixed;
    top: 155px;
    left: 50%;
    transform: translateX(-50%);
}
</style>