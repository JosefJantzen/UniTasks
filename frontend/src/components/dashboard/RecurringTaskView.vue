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
        <h1 style="text-align: center;">History</h1>
        <va-data-table :items="this.getHistory()" :columns="this.histCols">
            <template #header(count)>
                Counter
            </template>
            <template #cell(icon)="{ value }">
                <va-icon :name="value"/>
            </template>
        </va-data-table>
    </div>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

export default {
    name: "RecurringTaskView",
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
                task.history[i].count = parseInt(i) +1
                task.history[i].icon = this.getIcon(task.history[i].done)
            }
            console.log(task.history)
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
        }
    },
    data () {
        return {
            histCols: [
                {key: "count"},
                {key: "name"},
                {key: "icon"}
            ]
        }
    },
    props: {
        modal: Boolean,
        task: Object
    }
}

</script>