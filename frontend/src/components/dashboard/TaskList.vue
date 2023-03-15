<template>
    <va-list>
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
            <va-list-item-section icon>
                <va-button icon="check" round class="btn"/>
            </va-list-item-section>
            <va-list-item-section icon>
                <va-button color="light" icon="more_vert" round/>
            </va-list-item-section>
        </va-list-item>
    </va-list>
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
                    return  help.formatDate(task.ending)
                }
                return help.formatDate(task.due)
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

.btn {
    display: inline-block
}

</style>