<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['tasks/getAll'].filter(task => !task.done || showDone)"
        :key="index"
        :stripe="task.done ? true : false"
    >
        <div class="listItem" @click="show(task)">
            <va-card-title>
                <h1 style="font-size: 20px;">{{ task.name }}</h1>
                <div style="margin-left: auto; display: inline-block;" v-if="!task.done">
                    <va-icon name="mdi-schedule" /> 
                    <span style="margin-top: auto; margin-bottom: auto; margin-left: 0.5rem; font-size: small;"> {{ getDue(task)}} </span>
                </div>
                <va-button icon="mdi-check" round class="btn" style="margin-left: auto;" :disabled="task.done" @click="finished(task)"/>
                <va-button-dropdown
                    style="margin-left: 0.5rem;" 
                    preset="secondary" icon="more_vert" 
                    opened-icon="more_vert" 
                    round 
                    placement="right-start"
                    v-model="this.dropDown[index]"
                    @click.stop="this.dropDown[index] = !this.dropDown[index]"
                >
                    <va-button class="drop-btn" preset="secondary" icon="mdi-visibility" @click="show(task)">&nbsp;&nbsp;Show</va-button>
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
    <va-modal
        v-model="showModal"
        hide-default-actions
        size="medium"
    >
        <TaskView :modal="true" :task="this.modalTask" @click="close()"/>
    </va-modal>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

import TaskView from './TaskView.vue'

export default {
    name: 'OneTimeTasks',
    components: {
        TaskView
    },
    methods: {
        ...mapActions('tasks', ['getAll']),
        ...mapActions('tasks', ['done']),
        getDue (task) {
            return help.getDueString(task.due.substring(0,10))
        },
        undone (task) {
            task.done = false
            task.doneAt = null
            this.done(task)
        },
        finished (task) {
            task.done = true
            task.doneAt = help.now()
            this.done(task)
        },
        show (task) {
            this.showModal = true
            this.modalTask = task
        },
        close () {
            this.showModal = false
        }
    },
    data () {
        return {
            showModal: false,
            modalTask: null,
            dropDown: []
        }
    },
    props: {
        showDone: Boolean
    }
}

</script>