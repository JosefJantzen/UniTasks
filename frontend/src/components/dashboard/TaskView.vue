<template>
    <div class="" style="display: flex;">
            <va-button icon="mdi-close" size="small" round preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">{{ task.name }}</h1>
            <va-button icon="mdi-check" size="small" round v-if="!task.done" style="margin-left: auto;" @click="this.finish()"/>
            <va-button icon="mdi-edit" size="small" round preset="secondary" :style="task.done ? 'margin-left: auto;' : 'margin-left: 0.5rem;'" @click="showEdit(task)"/>
            <va-button icon="mdi-delete" size="small" round preset="secondary" style="margin-left: 0.5rem;" @click="this.delete(task)"/>
    </div>
    <br>
    <div style="margin-left: 1em; margin-right: 1em;">
        <p style="text-decoration: #99a4a5;">{{ task.desc }}</p><br>
        <div>
            <div style="display: flex;">
                <div style="display: flex; margin: auto 0; margin-right: 1em;">
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
    <va-modal
        v-model="showModalTaskEdit"
        hide-default-actions
        size="medium"
    >
        <TaskEdit :modal="true" :task="this.modalTask" :edit="true" @click="closeEdit()"/>
    </va-modal>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

import TaskEdit from '../TaskEdit.vue'

export default {
    name: "TaskView",
    components: {
        TaskEdit
    },
    methods: {
        ...mapActions('tasks', ['done']),
        ...mapActions('tasks', ['deleteTask']),
        ...mapActions('recurringTasks', ['doneHist']),
        ...mapActions('recurringTasks', ['deleteRecurringHist']),
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
        showEdit (task) {
            this.showModalTaskEdit = true
            this.modalTask = task            
        },
        closeEdit() {
            this.showModalTaskEdit = false
            this.showModalRecurringTaskEdit = false
        },
        delete (task) {
            this.$vaModal.init({
                title: 'Warning',
                message: 'Are you sure you want to delete this task?',
                okText: 'Yes',
                cancelText: 'No',
                blur: true,
                onOk: () => {
                    this.$emit('click')
                        if (task.recurring) {
                            this.deleteRecurringHist(task)
                        }
                        else {
                            this.deleteTask(task)
                        }
                },
            })            
        }
    },
    data () {
        return {
            showModalTaskEdit: false,
            modalTask: null
        }
    },
    props: {
        modal: Boolean,
        task: Object
    }
}

</script>