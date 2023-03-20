<template>
    <div :style="this.datePicker ? 'margin: 0 1em 5em' : 'margin: 0 1em 0'">
        <div class="" style="display: flex;">
            <va-button icon="mdi-close" size="small" round preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">New Task</h1>
        </div>
        <br>
        <va-form
            autofocus
            tag="form"
            @submit.prevent="submit"
        >
            <va-input
                v-model="name"
                label="Name"
                :rules="[(v) => v != '']"
                style="margin-bottom: 1em;"
                @click.stop=""
            /><br>
            <va-date-input
                v-model="due"
                v-model:is-open="datePicker"
                label="Due"
                first-weekday="Monday"
                style="margin-bottom: 1em;"
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
            <va-button type="submit">{{ getSubmitButton() }}</va-button>
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
        getSubmitButton () {
            if (this.edit) {
                return "Save"
            }
            return "Create"
        },
        async submit () {
            let task = this.task
            task.name = this.name
            task.due = help.formatJsDate(this.due) + "T23:59:59.999Z"
            task.desc = this.desc
            if (this.edit) {
                this.updateTask(task)
            } else {
                this.createTask(task)
            }
        }
    },
    data () {
        return {
            name: this.task.name,
            due: Date.now(),
            desc: this.task.desc,
            datePicker: false
        }
    },
    props: {
        task: Object,
        edit: Boolean
    }
}

</script>