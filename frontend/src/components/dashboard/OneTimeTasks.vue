<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['tasks/getAll'].filter(task => !task.done || showDone)"
        :key="index"
        :stripe="task.done ? true : false"
        :color="late(task) ? 'warning' : 'white'"
    >
        <div class="listItem" @click="showView(task)">
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
                    <va-button class="drop-btn" preset="secondary" icon="mdi-visibility" @click="showView(task)">&nbsp;&nbsp;Show</va-button>
                    <br>
                    <va-button class="drop-btn" preset="secondary" icon="mdi-edit" @click="showEdit(task)">&nbsp;&nbsp;&nbsp;Edit&nbsp;&nbsp;</va-button>
                    <br v-if="task.done">
                    <va-button class="drop-btn" preset="secondary" icon="mdi-undo" v-if="task.done" @click="undone(task)">Mark<br>undone</va-button>
                    <br>
                    <va-button class="drop-btn" preset="secondary" icon="mdi-delete" @click="this.delete(task)">Delete</va-button>
                </va-button-dropdown>
            </va-card-title>
        </div>
    </va-card>
    <va-modal
        v-model="showModalView"
        hide-default-actions
        size="medium"
    >
        <TaskView :modal="true" :task="this.modalTask" @click="closeView()"/>
    </va-modal>
    <va-modal
        v-model="showModalEdit"
        hide-default-actions
        size="medium"
    >
        <TaskEdit :modal="true" :task="this.modalTask" :edit="true" @click="closeEdit()"/>
    </va-modal>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

import TaskView from './TaskView.vue'
import TaskEdit from '../TaskEdit.vue'

export default {
    name: 'OneTimeTasks',
    components: {
        TaskView,
        TaskEdit
    },
    methods: {
        ...mapActions('tasks', ['getAll']),
        ...mapActions('tasks', ['done']),
        ...mapActions('tasks', ['deleteTask']),
        getDue (task) {
            return help.getDueString(task.due)
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
        showView (task) {
            this.showModalView = true
            this.modalTask = task
        },
        closeView () {
            this.showModalView = false
        },
        showEdit (task) {
            this.showModalEdit = true
            this.modalTask = task
        },
        closeEdit () {
            this.showModalEdit = false
        },
        delete (task) {
            this.$vaModal.init({
                title: 'Warning',
                message: 'Are you sure you want to delete this task?',
                okText: 'Yes',
                cancelText: 'No',
                blur: true,
                onOk: () => this.deleteTask(task),
            })
        },
        late (task) {
            return help.late(task.due)
        }
    },
    data () {
        return {
            showModalView: false,
            showModalEdit: false,
            modalTask: null,
            dropDown: []
        }
    },
    props: {
        showDone: Boolean
    }
}

</script>