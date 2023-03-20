<template>
    <div :style="this.datePicker ? 'margin: 0 1em 5em' : 'margin: 0 1em 0'">
        <div class="" style="display: flex;">
            <va-button icon="mdi-close" size="small" round preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">New Task</h1>
        </div>
        <br>
        <va-form
            autofocus
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

export default {
    name: "TaskEdit",
    methods: {
        getSubmitButton () {
            if (this.edit) {
                return "Save"
            }
            return "Create"
        }
    },
    data () {
        return {
            name: this.task.name,
            due: null,
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