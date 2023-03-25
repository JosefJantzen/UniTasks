<template>
    <div :style="this.startPicker || this.endPicker ? 'margin: 0 1em 5em' : 'margin: 0 1em 0'">
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
            /><br>
            <va-date-input
                v-model="start"
                v-model:is-open="startPicker"
                label="Start date"
                first-weekday="Monday"
                style="margin-bottom: 1em;"
                @click.stop="startPicker = !startPicker"
            /><br>
            <va-date-input
                v-model="ending"
                v-model:is-open="endPicker"
                label="End date"
                style="margin-bottom: 1em;"
                @click.stop="endPicker = !endPicker"
                :rules="[(ending > start || `Pick an end after the start`)]"
            /><br>
            <va-time-input
                v-model="due"
                v-model:is-open="duePicker"
                label="Due time of single tasks"
                style="margin-bottom: 1em;"
                @click.stop="duePicker = !duePicker"
            /><br>
            <va-input
                v-model="interval"
                label="Interval in days"
                type="number"
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
            <va-button type="submit" @click="this.$emit('click')" :disabled="ending < start || name == ''" >{{ getSubmitButton() }}</va-button>
        </va-form>
    </div>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../help/help'

export default {
    name: "RecurringTaskEdit",
    methods: {
        ...mapActions('recurringTasks', ['createRecurring']),
        ...mapActions('recurringTasks', ['createRecurringHist']),
        ...mapActions('recurringTasks', ['updateRecurring']),
        getHeading () {
            if (this.edit) {
                return "Edit Recurring Task"
            }
            return "New Recurring Task"
        },
        getSubmitButton () {
            if (this.edit) {
                return "Save"
            }
            return "Create"
        },
        async submit () {
            this.start.setHours(0, 0, 0, 0)
            this.ending.setHours(23, 59, 59, 999)
            let task = this.task
            task.name = this.name
            task.start = help.formatJsDate(this.start)
            task.ending = help.formatJsDate(this.ending)
            task.interval = this.interval
            task.desc = this.desc
            if (this.edit) {
                this.updateRecurring(task)
                
            } else {
                await this.createRecurring(task).then((recId) => {
                    this.start.setDate(this.start.getDate() + parseInt(this.interval)) 
                    while (this.start <= this.ending) {
                        this.createRecurringHist({
                            name: task.name,
                            desc: "",
                            due: help.formatJsDate(new Date(
                                this.start.getFullYear(), 
                                this.start.getMonth() + 1,
                                this.start.getDate(),
                                this.due.getHours(),
                                this.due.getMinutes(),
                                59,
                                999
                            )),
                            done: false,
                            recurringTaskId: recId
                        })
                        this.start.setDate(this.start.getDate() + parseInt(this.interval))  
                    }
                })    
            }
        }
    },
    data () {
        return {
            name: this.task.name,
            start: new Date(this.task.start),
            ending: new Date(this.task.ending),
            due: new Date(),
            desc: this.task.desc,
            interval: this.task.interval,
            startPicker: false,
            endPicker: false,
            duePicker: false
        }
    },
    props: {
        task: Object,
        edit: Boolean
    }
}

</script>