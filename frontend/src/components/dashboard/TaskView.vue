<template>
    <div class="" style="display: flex;">
            <va-button icon="mdi-close" size="small" round preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">{{ task.name }}</h1>
            <va-button icon="mdi-check" size="small" round v-if="!task.done" style="margin-left: auto;" @click="this.finish()"/>
            <va-button icon="mdi-edit" size="small" round preset="secondary" :style="task.done ? 'margin-left: auto;' : 'margin-left: 0.5rem;'"/>
            <va-button icon="mdi-delete" size="small" round preset="secondary" style="margin-left: 0.5rem;" />
    </div>
    <br>
    <div style="margin-left: 1em; margin-right: 1em;">
        <p style="text-decoration: #99a4a5;">{{ task.desc }}</p><br>
        <div>
            <div style="display: flex;">
                <div style="display: flex; margin: auto 0;">
                    <va-icon name="mdi-schedule"/>
                    <span style="margin-left: 1rem; margin-top: auto; margin-bottom: auto;">{{ this.getDueString() }}</span>
                </div>
                <va-avatar v-if="task.recurring" size="40px" font-size="15px" style="margin-left: auto;">{{ task.count }}/{{ task.countMax }}</va-avatar>
            </div>
            <div style="display: flex; margin-top: 1rem;">
                <va-icon v-if="task.done" name="mdi-check_circle"/>
                <span v-if="task.done" style="margin-left: 1rem; margin-top: auto; margin-bottom: auto;">{{ this.getDoneAtString() }}</span>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

export default {
    name: "TaskView",
    methods: {
        ...mapActions('tasks', ['done']),
        ...mapActions('recurringTasks', ['doneHist']),
        getDueString () {
            return "Due to " + help.formatTimestamp(this.task.due)
        },
        getDoneAtString () {
            return "Done at " + help.formatTimestamp(this.task.doneAt)
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
    },
    props: {
        modal: Boolean,
        task: Object
    }
}

</script>