<template>
    <div class="" style="display: flex;">
            <va-button icon="mdi-close" size="small" preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">{{ task.name }}</h1>
            <va-button icon="mdi-check" size="small" round v-if="!task.done" style="margin-left: auto;"/>
            <va-button icon="mdi-edit" size="small" preset="secondary" style="margin-left: 0.5rem;"/>
            <va-button icon="mdi-delete" size="small" preset="secondary" style="margin-left: 0.5rem;" />
    </div>
    <br>
    <div style="margin-left: 1em; margin-right: 1em;">
        <p style="text-decoration: #99a4a5;">{{ task.desc }}</p><br>
        <div style="display: flex;">
            <va-icon name="mdi-schedule"/>
            <span style="margin-left: 1rem; margin-top: auto; margin-bottom: auto;">{{ this.getDueString() }}</span>
        </div>
        <div v-if="task.done" style="display: flex; margin-top: 1rem;">
            <va-icon name="mdi-check_circle"/>
            <span style="margin-left: 1rem; margin-top: auto; margin-bottom: auto;">{{ this.getDoneAtString() }}</span>
        </div>
    </div>
</template>

<script>
import help from '../help/help'

export default {
    name: "TaskView",
    methods: {
        getDueString () {
            return "Due to " + help.formatTimestamp(this.task.due)
        },
        getDoneAtString () {
            return "Done at " + help.formatTimestamp(this.task.doneAt)
        }
    },
    props: {
        modal: Boolean,
        task: Object
    }
}

</script>