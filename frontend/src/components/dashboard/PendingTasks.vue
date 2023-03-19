<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['getPendingTasks'].filter(task => !task.done || showDone)"
        :key="index"
        :stripe="task.done ? true : false"
    >
        <div class="listItem">
            <va-card-title>
                <va-avatar v-if="task.recurring" size="40px" font-size="15px">{{ task.count }}/{{ task.countMax }}</va-avatar>
                <va-avatar v-else icon="mdi-repeat_one" size="40px"/>
                <h1 style="font-size: 20px; margin-left: 0.5rem;">{{ task.name }}</h1>
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
            <va-card-content style="display: flex;" v-if="!task.done">
                <div style="display: flex;" >
                    <va-icon name="mdi-schedule" style="margin-top: auto; margin-bottom: auto;"/> 
                    <span style="margin-top: auto; margin-bottom: auto; margin-left: 0.5rem;"> {{ getDue(task)}} </span>
                </div>
            </va-card-content>
        </div>
    </va-card>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

export default {
    name: 'PendingTasks',
    methods: {
        ...mapActions('tasks', ['list']),
        ...mapActions('tasks', ['done']),
        ...mapActions('recurringTasks', ['listRecurring']),
        ...mapActions('recurringTasks', ['doneHist']),
        getDue (task) {
            return help.getDueString(task.due.substring(0,10))

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
        }
    },
    created () {
        this.list()
        this.listRecurring()
    },
    data () {
        return {
            hoverItem: false
        }
    },
    props: {
        showDone: Boolean
    }
}

</script>

<style>

.task-done {
    background-color: #000000;
}

.listItem{
  margin-top: 20px;
  cursor: pointer;
}

.listItem:hover {
    background-color: #d5e8e8;
}

.card-content {
    padding: 1rem;
    overflow: hidden;
}

.column {
    display: contents;
    margin-right: 1rem;
    float: left;
}

.drop-btn {
    margin-left: 1rem;
    margin-right: 1rem;
}

</style>