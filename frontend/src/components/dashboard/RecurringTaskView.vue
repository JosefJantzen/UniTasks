<template>
    <div class="" style="display: flex;">
            <va-button icon="mdi-close" size="small" round preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">{{ task.name }}</h1>
            <va-button icon="mdi-edit" size="small" round preset="secondary" style="margin-left: auto;"/>
            <va-button icon="mdi-delete" size="small" round preset="secondary" style="margin-left: 0.5rem;" />
    </div>
    <br>
    <div style="margin-left: 1em; margin-right: 1em;">
        <p style="text-decoration: #99a4a5;">{{ task.desc }}</p><br>
        <div style="display: flex;">
            <va-icon name="mdi-schedule"/>
            <span style="margin-left: 0.5rem; margin-top: auto; margin-bottom: auto;">{{ this.getTimeString() }} every {{ task.interval }} days</span>
        </div>
        <br>
        <span><b>Next date:&nbsp;</b> {{ getNextDueString() }}</span>
    </div>
    <br>
    <div>
        <h1 style="text-align: center; margin-bottom: 0.25rem;">History</h1>
        <va-data-table 
            :items="this.getHistory()" 
            :columns="this.histCols"
            clickable hoverable
            sticky-header
            :item-size="46"
            :wrapper-size="400"
            @row:click="rowClick"
        >
            <template #header(count)>
                Counter
            </template>
            <template #header(icon)>
                Done
            </template>
            <template #header(actions)>
                {{ this.countDoneHist() }}/{{ this.task.history.length }}
            </template>
            <template #cell(count)="{ rowIndex }">
                {{ rowIndex +1 }}
            </template>
            <template #cell(icon)="{ value }">
                <va-icon :name="value"/>
            </template>
            <template #cell(actions)="{ rowIndex }">
                <va-button preset="secondary" round icon="mdi-edit" @click.stop="rowIndex"/>
                <va-button preset="secondary" round icon="mdi-delete" @click.stop="rowIndex"/>
            </template>
        </va-data-table>
    </div>
    <va-modal
        v-model="showModal"
        hide-default-actions
        size="medium"
        blur
    >
        <TaskView :modal="true" :task="this.modalTask" @click="close()"/>
    </va-modal>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

import TaskView from './TaskView.vue'

export default {
    name: "RecurringTaskView",
    components: {
        TaskView
    },
    methods: {
        ...mapActions('tasks', ['done']),
        ...mapActions('recurringTasks', ['doneHist']),
        getTimeString () {
            return "From " + help.formatDate(this.task.start) + " to " + help.formatDate(this.task.ending)
        },
        getDoneAtString () {
            return "Done at " + help.formatTimestamp(this.task.doneAt)
        },
        getNextDueString () {
            return help.formatTimestamp(this.getNextHist().due)
        },
        getHistory () {
            let task = this.task
            for(const i in task.history) {
                task.history[i].icon = this.getIcon(task.history[i].done)
            }
            return task.history
        },
        getNextHist() {
            let task = this.task
            for(const i in task.history) {
                if(!task.history[i].done) {
                    return task.history[i]
                }
            }
        },
        countDoneHist() {
            let res = 0
            for (const t of this.task.history) {
                if (t.done) {
                    res++
                }
            }
            return res
        },
        finish () {
            let task = this.task
            task.done = true
            task.doneAt = help.now()
            if (task.recurring) {
                this.doneHist(task)
                return
            }
            this.done(task)
        },
        getIcon (v) {
            if (v) {
                return "mdi-check"
            }
            return "mdi-close"
        },
        rowClick(event) {
            this.modalTask = event.item
            this.showModal = true
        },
        show (task) {
            this.modalTask = task
        },
        close () {
            this.showModal = false
        }
    },
    data () {
        return {
            histCols: [
                {key: "count"},
                {key: "name"},
                {key: "icon"},
                {key: "actions"}
            ],
            showModal: false,
            modalTask: null,
        }
    },
    props: {
        modal: Boolean,
        task: Object
    }
}

</script>