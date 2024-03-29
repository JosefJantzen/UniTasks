<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['getPendingTasks'].filter(task => !task.done || showDone)"
        :key="index"
        :stripe="task.done || late(task) ? true : false"
        :stripe-color="late(task) && !task.done ? 'warning' : 'primary'"
    >
        <div class="listItemPending" @click="show(task)">
            <va-card-title>
                <va-avatar v-if="task.recurring" size="40px" font-size="15px">{{ task.count }}/{{ task.countMax }}</va-avatar>
                <va-avatar v-else icon="mdi-repeat_one" size="40px"/>
                <h1 style="font-size: 20px; margin-left: 0.5rem;">{{ task.name }}</h1>
                <va-button icon="mdi-check" round class="btn" style="margin-left: auto;" :disabled="task.done" @click.stop="finished(task)"/>
                <va-button-dropdown 
                    style="margin-left: 0.5rem;" 
                    preset="secondary" icon="more_vert" 
                    opened-icon="more_vert" 
                    round 
                    placement="right-start"
                    v-model="this.dropDown[index]"
                    @click.stop="this.dropDown[index] = !this.dropDown[index]"
                >
                    <va-button class="dropBtnPending" preset="secondary" icon="mdi-visibility" @click="show(task)">&nbsp;&nbsp;Show</va-button>
                    <br>
                    <va-button class="dropBtnPending" preset="secondary" icon="mdi-edit" @click="showEdit(task)">&nbsp;&nbsp;&nbsp;Edit&nbsp;&nbsp;</va-button>
                    <br v-if="task.done">
                    <va-button class="dropBtnPending" preset="secondary" icon="mdi-undo" v-if="task.done" @click="undone(task)">Mark<br>undone</va-button>
                    <br>
                    <va-button class="dropBtnPending" preset="secondary" icon="mdi-delete" @click="this.delete(task)">Delete</va-button>
                </va-button-dropdown>
            </va-card-title>
            <va-card-content style="display: flex;" v-if="!task.done">
                <div style="display: flex;" >
                    <va-icon name="mdi-schedule" style="margin-top: auto; margin-bottom: auto;"/> 
                    <span style="margin-top: auto; margin-bottom: auto; margin-left: 0.5rem;"> {{ getDue(task)}} </span>
                </div>
            </va-card-content>
        </div>
    </va-card>
    <va-modal
        v-model="showModal"
        hide-default-actions
        size="medium"
    >
        <TaskView :modal="true" :task="this.modalTask" @click="close()"/>
    </va-modal>
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

import TaskView from './TaskView.vue'
import TaskEdit from '../TaskEdit.vue'

export default {
    name: 'PendingTasks',
    components: {
        TaskView,
        TaskEdit
    },
    methods: {
        ...mapActions('tasks', ['listTask']),
        ...mapActions('tasks', ['done']),
        ...mapActions('tasks', ['deleteTask']),
        ...mapActions('recurringTasks', ['listRecurring']),
        ...mapActions('recurringTasks', ['doneHist']),
        ...mapActions('recurringTasks', ['deleteRecurringHist']),
        getDue (task) {
            return help.getDueString(task.due)

        },
        undone (task) {
            task.done = false
            task.doneAt = null
            if (task.recurring) {
                this.doneHist(task)
                return
            }
            this.done(task)
        },
        finished (task) {
            task.done = true
            task.doneAt = help.now()
            if (task.recurring) {
                this.doneHist(task)
                return
            }
            this.done(task)
        },
        show (task) {
            this.showModal = true
            this.modalTask = task
        },
        close () {
            this.showModal = false
        },
        showEdit (task) {
            this.showModalTaskEdit = true
            this.modalTask = task
        },
        closeEdit() {
            this.showModalTaskEdit = false
        },
        delete (task) {
            this.$vaModal.init({
                title: 'Warning',
                message: 'Are you sure you want to delete this task?',
                okText: 'Yes',
                cancelText: 'No',
                blur: true,
                onOk: () => {
                    if (task.recurring) {
                        this.deleteRecurringHist(task)
                    }
                    else {
                        this.deleteTask(task)
                    }
                },
            })
        },
        late (task) {
            return help.late(task.due)
        }
    },
    created () {
        this.listTask()
        this.listRecurring()
    },
    data () {
        return {
            showModal: false,
            modalTask: null,
            dropDown: [],
            showModalTaskEdit: false,
        }
    },
    props: {
        showDone: Boolean
    }
}

</script>

<style>

.listItemPending {
  margin-top: 20px;
  cursor: pointer;
}

.listItemPending:hover {
    background-color: #d5e8e8;
}

.dropBtnPending {
    margin-left: 1rem;
    margin-right: 1rem;
}

</style>