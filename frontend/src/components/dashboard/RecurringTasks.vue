<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['recurringTasks/getAll'].filter(task => this.filter(task) || showDone)"
        :key="index"
        :stripe="!this.filter(task)"
    >
        <div class="listItem" @click="show(task)">
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
                    <va-button class="drop-btn" preset="secondary" icon="mdi-visibility"  @click="show(task)">&nbsp;&nbsp;Show</va-button>
                    <br>
                    <va-button class="drop-btn" preset="secondary" icon="mdi-edit">&nbsp;&nbsp;&nbsp;Edit&nbsp;&nbsp;</va-button>
                    <br v-if="this.filter(task)">
                    <va-button class="drop-btn" preset="secondary" icon="mdi-cancel" v-if="this.filter(task)" @click="end(task)">&nbsp;End it</va-button>
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
        <RecurringTaskView :modal="true" :task="this.modalTask" @click="close()"/>
    </va-modal>
</template>

<script>
import moment from 'moment'
import { mapActions } from 'vuex'
import help from '../../help/help'

import RecurringTaskView from './RecurringTaskView.vue'

export default {
    name: 'RecurringTasks',
    components: {
        RecurringTaskView
    },
    methods: {
        ...mapActions('recurringTasks', ['updateRecurring']),
        getDue (task) {
            return help.getDueString(task.ending.substring(0,10))
        },
        end (task) {
            task.ending = help.now()
            this.updateRecurring(task)
        },
        filter (task) {
            return moment.utc(task.ending).isAfter(moment.utc());
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