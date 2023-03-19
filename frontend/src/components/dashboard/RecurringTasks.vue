<template>
    <va-card 
        v-for="(task, index) in this.$store.getters['recurringTasks/getAll'].filter(task => this.filter(task) || showDone)"
        :key="index"
        :stripe="!this.filter(task)"
    >
        <div class="listItem">
            <va-card-title>
                <h1 style="font-size: 20px;">{{ task.name }}</h1>
                <div style="margin-left: auto; display: inline-block;" v-if="this.filter(task)">
                    <va-icon name="schedule" /> 
                    <span style="margin-top: auto; margin-bottom: auto; margin-left: 0.5rem; font-size: small;"> {{ getDue(task)}} </span>
                </div>
                <va-button-dropdown 
                    style="margin-left: auto;" 
                    preset="plain" icon="more_vert" 
                    opened-icon="more_vert" 
                    round 
                    placement="right-start"
                >
                    <va-button class="drop-btn" preset="secondary" icon="mdi-visibility">&nbsp;&nbsp;Show</va-button>
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
</template>

<script>
import moment from 'moment'
import { mapActions } from 'vuex'
import help from '../../help/help'

export default {
    name: 'RecurringTasks',
    methods: {
        ...mapActions('recurringTasks', ['update']),
        getDue (task) {
            return help.getDueString(task.ending.substring(0,10))
        },
        end (task) {
            task.ending = help.now()
            this.update(task)
        },
        filter (task) {
            return moment.utc(task.ending).isAfter(moment.utc());
        }
    },
    props: {
        showDone: Boolean
    }
}

</script>