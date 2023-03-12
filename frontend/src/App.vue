<template #app>
	<va-navbar color="primary" class="navbar">
		<template #left>
			<va-navbar-item class="nav-logo">
				<va-image src="logo.png" :max-width=40 style="cursor: pointer;" @click="$router.push('/')"/>
			</va-navbar-item>
		</template>
		<template #center>
			<va-navbar-item>All Tasks</va-navbar-item>
			<va-navbar-item>Tasks</va-navbar-item>
			<va-navbar-item>Recurring Tasks</va-navbar-item>
		</template>
		<template #right>
			<AccountMenu></AccountMenu>
		</template>
	</va-navbar>
	<router-view></router-view>
</template>

<script>
import AccountMenu from "./components/navbar/AccountMenu.vue";

export default {
	name: 'App',
	components: {
		AccountMenu
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
</style>
