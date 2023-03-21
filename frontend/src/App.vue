<template #app>
	<va-navbar color="primary" class="navbar">
		<template #left>
			<va-navbar-item class="nav-logo">
				<va-image src="logo.png" :max-width=40 style="cursor: pointer;" @click="$router.push('/')"/>
			</va-navbar-item>
		</template>
		<template #center v-if="this.$store.getters['user/get'] != null">
			<va-navbar-item>
				<va-button size="large" @click="$router.replace('/')">Dashboard</va-button>
			</va-navbar-item>
			<va-navbar-item>
				<va-button size="large" @click="showNewTask()">New Task</va-button>
			</va-navbar-item>
			<va-navbar-item>
				<va-button size="large" @click="showNewRecurringTask()">New Recurring Task</va-button>
			</va-navbar-item>
		</template>
		<template #right>
			<va-button class="git" @click="github">GitHub</va-button>
			<AccountMenu></AccountMenu>
		</template>
	</va-navbar>
	<router-view></router-view>
	<va-modal
        v-model="showModalNewTask"
        hide-default-actions
        size="large"
    >
        <TaskEdit @click="closeNewTask()" :task="createEmptyTask()" :edit="false"/>
    </va-modal>
	<va-modal
        v-model="showModalNewRecurringTask"
        hide-default-actions
        size="large"
    >
        <RecurringTaskEdit @click="closeNewRecurringTask()" :task="createEmptyRecurringTask()" :edit="false"/>
    </va-modal>
</template>

<script>
import AccountMenu from "./components/navbar/AccountMenu.vue";
import TaskEdit from "./components/TaskEdit.vue";
import RecurringTaskEdit from "./components/RecurringTaskEdit.vue";

export default {
	name: 'App',
	components: {
		AccountMenu,
		TaskEdit,
		RecurringTaskEdit
	},
	methods: {
		github () {
			window.open("https://github.com/JosefJantzen/UniTasks", '_blank');
		},
		showNewTask () {
            this.showModalNewTask = true
        },
        closeNewTask() {
            this.showModalNewTask = false
        },
		showNewRecurringTask () {
            this.showModalNewRecurringTask = true
        },
        closeNewRecurringTask() {
            this.showModalNewRecurringTask = false
        },
		createEmptyTask () {
			return {
				name: "",
				desc: "",
				due: new Date().toISOString(),
				done: false,
				doneAt: null
			}
		},
		createEmptyRecurringTask () {
			let d = new Date()
			d.setDate(d.getDate() + 1)
			return {
				name: "",
				desc: "",
				start: new Date().toISOString(),
				ending: d.toISOString(),
				interval: 7,
				history: []
			}
		}
	},
	data () {
		return {
			showModalNewTask: false,
			showModalNewRecurringTask: false
		}
	}
}
</script>

<style>
#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
}

.navbar {
	margin-bottom: 1em;
	box-shadow: 0 2px 8px rgba(0,0,0,.5);
}

.git {
	margin-right: 1rem;
}
</style>
