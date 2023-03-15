<template>
    <!--<va-list>
        <va-list-item
            v-for="(task, index) in this.$store.getters['getAllTasks']"
            :key="index"
            class="list__item"
        >   
            <va-list-item-section avatar>
                <va-avatar v-if="task.recurring" icon="mdi-repeat" />
                <va-avatar v-else icon="mdi-repeat_one" />
            </va-list-item-section>
            <va-list-item-section>
                <va-list-item-label>
                    {{ task.name }} 
                </va-list-item-label>
                <va-list-item-label caption>
                    {{ task.desc }}
                </va-list-item-label>
            </va-list-item-section>
            <va-list-item-section>
                <div>
                    <va-icon name="schedule"/> {{ getDue(task)}}
                </div>
            </va-list-item-section>
            <va-list-separator></va-list-separator>
            <va-list-item-section avatar>
                <va-button icon="check" round class="btn"/>
            </va-list-item-section>
            <va-list-item-section avatar>
                <va-button-dropdown color="light" icon="more_vert" opened-icon="more_vert" round placement="right-start">
                    <va-button class="drop-btn" preset="secondary" icon="visibility">&nbsp;&nbsp;Show</va-button>
                    <br>
                    <va-button class="drop-btn" preset="secondary" icon="edit">&nbsp;&nbsp;&nbsp;Edit&nbsp;&nbsp;</va-button>
                    <br>
                    <va-button class="drop-btn" preset="secondary" icon="delete">Delete</va-button>
                </va-button-dropdown>
            </va-list-item-section>
        </va-list-item>
    </va-list>-->
    <va-card 
        v-for="(task, index) in this.$store.getters['getAllTasks']"
        :key="index"
        class="list__item"
    >
        <va-card-title>
            <va-avatar v-if="task.recurring" icon="mdi-repeat" size="40px"/>
            <va-avatar v-else icon="mdi-repeat_one" size="40px"/>
            <h1 style="font-size: 20px; margin-left: 0.5rem;">{{ task.name }}</h1>
            <va-button icon="mdi-check" round class="btn" style="margin-left: auto;"/>
            <va-button-dropdown style="margin-left: 0.5rem;" preset="plain" icon="more_vert" opened-icon="more_vert" round placement="right-start">
                <va-button class="drop-btn" preset="secondary" icon="visibility">&nbsp;&nbsp;Show</va-button>
                <br>
                <va-button class="drop-btn" preset="secondary" icon="edit">&nbsp;&nbsp;&nbsp;Edit&nbsp;&nbsp;</va-button>
                <br>
                <va-button class="drop-btn" preset="secondary" icon="delete">Delete</va-button>
            </va-button-dropdown>
        </va-card-title>
        <va-card-content style="display: flex;">
            <va-icon name="schedule"/> 
            <p style="margin-top: auto; margin-bottom: auto; margin-left: 0.5rem;"> {{ getDue(task)}} </p>
        </va-card-content>
    </va-card>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

export default {
    name: 'TaskList',
    methods: {
        ...mapActions('tasks', ['list']),
        ...mapActions('recurringTasks', ['listRecurring']),
        getDue (task) {
            if (task.recurring) {
                    return  help.getDueString(task.ending.substring(0,10))
                }
                return help.getDueString(task.due.substring(0,10))
        }
    },
    created () {
        this.list()
        this.listRecurring()
    }
}

</script>

<style>

.list__item + .list__item {
  margin-top: 20px;
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