<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['recurringTasks/getAll'].filter(task => this.filter(task) || showDone)"
        :key="index"
        :stripe="!this.filter(task)"
    >
        <div class="listItemRecurring" @click="show(task)">
            <va-card-title>
                <h1 style="font-size: 20px;">{{ task.name }}</h1>
                <div style="margin-left: auto; display: inline-block;" v-if="this.filter(task)">
                    <va-icon name="schedule" /> 
                    <span style="margin-top: auto; margin-bottom: auto; margin-left: 0.5rem; font-size: small;"> {{ getDue(task)}} </span>
                </div>
                <va-button-dropdown 
                    style="margin-left: auto;" 
                    preset="secondary" icon="more_vert" 
                    opened-icon="more_vert" 
                    round 
                    placement="right-start"
                    v-model="this.dropDown[index]"
                    @click.stop="this.dropDown[index] = !this.dropDown[index]"
                >
                    <va-button class="dropBtnRecurring" preset="secondary" icon="mdi-visibility"  @click="show(task)">&nbsp;&nbsp;Show</va-button>
                    <br>
                    <va-button class="dropBtnRecurring" preset="secondary" icon="mdi-edit" @click="showEdit(task)">&nbsp;&nbsp;&nbsp;Edit&nbsp;&nbsp;</va-button>
                    <br v-if="this.filter(task)">
                    <va-button class="dropBtnRecurring" preset="secondary" icon="mdi-cancel" v-if="this.filter(task)" @click="end(task)">&nbsp;End it</va-button>
                    <br>
                    <va-button class="dropBtnRecurring" preset="secondary" icon="mdi-delete" @click="this.delete(task)">Delete</va-button>
                </va-button-dropdown>
            </va-card-title>
        </div>
    </va-card>
    <va-modal
        v-model="showModal"
        hide-default-actions
        size="medium"
    >
        <RecurringTaskView :modal="true" :task="this.modalTask" @click="close()"/>
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
import moment from 'moment'
import { mapActions } from 'vuex'
import help from '../../help/help'

import RecurringTaskView from './RecurringTaskView.vue'
import RecurringTaskEdit from '../RecurringTaskEdit.vue'

export default {
    name: 'RecurringTasks',
    components: {
        RecurringTaskView,
        RecurringTaskEdit
    },
    methods: {
        ...mapActions('recurringTasks', ['updateRecurring']),
        ...mapActions('recurringTasks', ['deleteRecurring']),
        getDue (task) {
            return help.getDueString(task.ending)
        },
        end (task) {
            task.ending = help.now()
            this.updateRecurring(task)
        },
        filter (task) {
            return moment.utc(task.ending).isAfter(moment.utc());
        },
        show (task) {
            this.modalTask = task
            this.showModal = true
        },
        close () {
            this.showModal = false
        },
        showEdit (task) {
            this.modalTaskEdit = task
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
                onOk: () => this.deleteRecurring(task),
            })            
        }
    },
    data () {
        return {
            showModal: false,
            showModalEdit: false,
            modalTask: null,
            modalTaskEdit: null,
            dropDown: []
        }
    },
    props: {
        showDone: Boolean
    }
}

</script>

<style>

.listItemRecurring{
  margin-top: 20px;
  cursor: pointer;
}

.listItemRecurring:hover {
    background-color: #d5e8e8;
}

.dropBtnRecurring {
    margin-left: 1rem;
    margin-right: 1rem;
}

</style>