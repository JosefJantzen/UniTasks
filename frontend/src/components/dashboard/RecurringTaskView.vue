<template>
    <div class="" style="display: flex;">
            <va-button icon="mdi-close" size="small" round preset="secondary" @click="this.$emit('click')"/>
            <h1 style="margin: auto 1rem; font-size: 25px;">{{ task.name }}</h1>
            <va-button icon="mdi-edit" size="small" round preset="secondary" style="margin-left: auto;" @click="showEdit()"/>
            <va-button icon="mdi-delete" size="small" round preset="secondary" style="margin-left: 0.5rem;" @click="this.delete(task)"/>
    </div>
    <br>
    <div style="margin-left: 1em; margin-right: 1em;">
        <p style="text-decoration: #99a4a5;">{{ task.desc }}</p><br>
        <div style="display: flex;">
            <va-icon name="mdi-schedule"/>
            <span style="margin-left: 0.5rem; margin-top: auto; margin-bottom: auto;">{{ this.getTimeString() }} every {{ task.interval }} days</span>
        </div>
        <br>
        <span><b>Next date:&nbsp;</b> {{ getNextDueString() }}</span>
    </div>
    <br>
    <div>
        <div style="display: flex; margin-bottom: 0.25rem;">
            <h1 style="text-align: center; margin: auto 0;">History</h1>
            <va-button icon="mdi-add" preset="secondary" size="small" style="margin: auto 0 auto auto;" @click="showNew()"/>
        </div>
        <va-data-table 
            :items="this.getHistory()" 
            :columns="this.histCols"
            clickable hoverable
            sticky-header
            :item-size="46"
            :wrapper-size="400"
            @row:click="rowClick"
            virtual-scroller
        >
            <template #header(count)>
                Counter
            </template>
            <template #header(icon)>
                Done
            </template>
            <template #header(actions)>
                {{ this.countDoneHist() }}/{{ this.task.history.length }}
            </template>
            <template #cell(count)="{ rowIndex }">
                {{ rowIndex +1 }}
            </template>
            <template #cell(icon)="{ value }">
                <va-icon :name="value"/>
            </template>
            <template #cell(actions)="{ rowIndex }">
                <va-button style="z-index: 0;" preset="secondary" round icon="mdi-edit" @click.stop="this.showEditHist(rowIndex)"/>
                <va-button style="z-index: 0;" preset="secondary" round icon="mdi-delete" @click.stop="this.deleteHist(rowIndex)"/>
            </template>
        </va-data-table>
    </div>
    <va-modal
        v-model="showModal"
        hide-default-actions
        size="medium"
        blur
    >
        <TaskView :modal="true" :task="this.modalRec" @click="close()"/>
    </va-modal>
    <va-modal
        v-model="showModalTaskNew"
        hide-default-actions
        size="medium"
    >
        <TaskEdit :modal="true" :task="this.createEmptyTaskHist()" :edit="false" @click="closeNew()"/>
    </va-modal>
    <va-modal
        v-model="showModalTaskEdit"
        hide-default-actions
        size="medium"
    >
        <TaskEdit :modal="true" :task="this.modalHist" :edit="true" @click="closeEditHist()"/>
    </va-modal>
    <va-modal
        v-model="showModalEdit"
        hide-default-actions
        size="medium"
    >
        <RecurringTaskEdit :modal="true" :edit="true" :task="this.modalTaskEdit" @click="closeEdit()"/>
    </va-modal>
</template>


<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

import TaskView from './TaskView.vue'
import TaskEdit from '../TaskEdit.vue'
import RecurringTaskEdit from '../RecurringTaskEdit.vue'

export default {
    name: "RecurringTaskView",
    components: {
        TaskView,
        TaskEdit,
        RecurringTaskEdit
    },
    methods: {
        ...mapActions('tasks', ['done']),
        ...mapActions('recurringTasks', ['doneHist']),
        ...mapActions('recurringTasks', ['deleteRecurring']),
        ...mapActions('recurringTasks', ['deleteRecurringHist']),
        getTimeString () {
            return "From " + help.formatDate(this.task.start) + " to " + help.formatDate(this.task.ending)
        },
        getDoneAtString () {
            return "Done at " + help.formatTimestamp(this.task.doneAt)
        },
        getNextDueString () {
            let h = this.getNextHist()
            if (h == null) {
                return "No entries"
            }
            return help.formatTimestamp(h.due)
        },
        getHistory () {
            let task = this.task
            for(const i in task.history) {
                task.history[i].icon = this.getIcon(task.history[i].done)
            }
            return task.history
        },
        getNextHist() {
            let task = this.task
            for(const i in task.history) {
                if(!task.history[i].done) {
                    return task.history[i]
                }
            }
        },
        countDoneHist() {
            let res = 0
            for (const t of this.task.history) {
                if (t.done) {
                    res++
                }
            }
            return res
        },
        finish () {
            let task = this.task
            task.done = true
            task.doneAt = help.now()
            if (task.recurring) {
                this.doneHist(task)
                return
            }
            this.done(task)
        },
        getIcon (v) {
            if (v) {
                return "mdi-check"
            }
            return "mdi-close"
        },
        rowClick(event) {
            this.modalRec= event.item
            this.showModal = true
        },
        close () {
            this.modalRec = false
        },
        showNew () {
            this.showModalTaskNew = true
        },
        closeNew () {
            this.showModalTaskNew = false
        },
        showEditHist (i) {
            this.modalHist = this.getHistory()[i]
            this.showModalTaskEdit = true
        },
        closeEditHist () {
            this.showModalTaskEdit = false
        },
        showEdit () {
            this.modalTaskEdit = this.task
            this.showModalEdit = true
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
                onOk: () => {
                    this.$emit('click')
                    this.deleteRecurring(task)
                },
            })            
        },
        deleteHist (i) {
            this.$vaModal.init({
                title: 'Warning',
                message: 'Are you sure you want to delete this task?',
                okText: 'Yes',
                cancelText: 'No',
                blur: true,
                onOk: () => this.deleteRecurringHist(this.getHistory()[i]),
            })            
        },
		createEmptyTaskHist () {
			return {
				name: this.task.name,
				desc: "",
				due: new Date().toISOString(),
				done: false,
				doneAt: null,
                recurring: true,
                recurringTaskId: this.task.id
			}
		}
    },
    data () {
        return {
            histCols: [
                {key: "count"},
                {key: "name"},
                {key: "icon"},
                {key: "actions"}
            ],
            showModal: false,
            showModalTaskNew: false,
            showModalTaskEdit: false,
            showModalEdit: false,
            modalTask: null,
            modalHist: null,
            modalTaskEdit: null
        }
    },
    props: {
        modal: Boolean,
        task: Object
    }
}

</script>