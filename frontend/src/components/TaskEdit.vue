<template>
    <div :style="this.datePicker || this.timePicker ? 'margin: 0 1em 5em' : 'margin: 0 1em 0'">
        <div class="" style="display: flex;" @click.stop="">
            <va-button icon="mdi-close" size="small" round preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">{{ getHeading() }}</h1>
        </div>
        <br>
        <va-form
            autofocus
            tag="form"
            @submit.prevent="submit"
            @click.stop=""
        >
            <va-input
                v-model="name"
                label="Name"
                :rules="[(v) => v != '']"
                style="margin-bottom: 1em;"
                @click.stop=""
                :disabled="this.task.recurring"
            /><br>
            <va-date-input
                v-model="due"
                v-model:is-open="datePicker"
                label="Due date"
                first-weekday="Monday"
                style="margin-bottom: 1em;"
                @click.stop="datePicker = !datePicker"
            /><br>
            <va-time-input
                v-model="dueTime"
                v-model:is-open="timePicker"
                label="Due time"
                style="margin-bottom: 1em;"
                @click.stop="timePicker = !timePicker"
            /><br>
            <va-input
                v-model="desc"
                label="Description"
                type="textarea"
                :min-rows="3"
                :max-rows="5"
                style="margin-bottom: 1.5em;"
                @click.stop=""
            >
            </va-input><br>
            <va-button type="submit" @click="this.$emit('click')" :disabled="name == ''">{{ getSubmitButton() }}</va-button>
        </va-form>
    </div>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../help/help'

export default {
    name: "TaskEdit",
    methods: {
        ...mapActions('tasks', ['createTask']),
        ...mapActions('tasks', ['updateTask']),
        ...mapActions('recurringTasks', ['updateRecurringHist']),
        ...mapActions('recurringTasks', ['createRecurringHist']),
        getHeading () {
            if (this.edit) {
                return "Edit Task"
            }
            return "New Task"
        },
        getSubmitButton () {
            if (this.edit) {
                return "Save"
            }
            return "Create"
        },
        async submit () {
            this.due.setHours(
                this.dueTime.getHours(),
                this.dueTime.getMinutes(),
                59,
                999
            )
            let task = this.task
            task.name = this.name
            task.due = help.formatJsDate(this.due)
            task.desc = this.desc
            if (this.edit) {
                if (task.recurring) {
                    this.updateRecurringHist(task)
                } else {
                    this.updateTask(task)
                }
                
            } else {
                if (task.recurring) {
                    this.createRecurringHist(task)
                } else {
                    this.createTask(task)                    
                }
            }
        }
    },
    data () {
        return {
            name: this.task.name,
            due: new Date(this.task.due),
            dueTime: new Date(this.task.due),
            desc: this.task.desc,
            datePicker: false,
            timePicker: false
        }
    },
    props: {
        task: Object,
        edit: Boolean
    }
}

</script>