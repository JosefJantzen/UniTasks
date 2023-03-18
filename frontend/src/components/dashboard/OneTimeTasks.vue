<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['tasks/getAll'].filter(task => !task.done || showDone)"
        :key="index"
        :stripe="task.done ? true : false"
    >
        <div class="listItem">
            <va-card-title>
                <h1 style="font-size: 20px;">{{ task.name }}</h1>
                <div style="margin-left: auto; display: inline-block;" v-if="!task.done">
                    <va-icon name="schedule" /> 
                    <span style="margin-top: auto; margin-bottom: auto; margin-left: 0.5rem; font-size: small;"> {{ getDue(task)}} </span>
                </div>
                <va-button icon="mdi-check" round class="btn" style="margin-left: auto;" :disabled="task.done" @click="finished(task)"/>
                <va-button-dropdown
                    style="margin-left: 0.5rem;" 
                    preset="plain" icon="more_vert" 
                    opened-icon="more_vert" 
                    round 
                    placement="right-start"
                >
                    <va-button class="drop-btn" preset="secondary" icon="mdi-visibility">&nbsp;&nbsp;Show</va-button>
                    <br>
                    <va-button class="drop-btn" preset="secondary" icon="mdi-edit">&nbsp;&nbsp;&nbsp;Edit&nbsp;&nbsp;</va-button>
                    <br v-if="task.done">
                    <va-button class="drop-btn" preset="secondary" icon="mdi-undo" v-if="task.done" @click="undone(task)">Mark<br>undone</va-button>
                    <br>
                    <va-button class="drop-btn" preset="secondary" icon="mdi-delete">Delete</va-button>
                </va-button-dropdown>
            </va-card-title>
        </div>
    </va-card>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

export default {
    name: 'OneTimeTasks',
    methods: {
        ...mapActions('tasks', ['getAll']),
        ...mapActions('tasks', ['done']),
        getDue (task) {
            return help.getDueString(task.due.substring(0,10))
        },
        undone (task) {
            task.done = false
            task.doneAt = ""
            this.done(task)
        },
        finished (task) {
            task.done = true
            task.doneAt = help.now()
            this.done(task)
        }
    },
    props: {
        showDone: Boolean
    }
}

</script>