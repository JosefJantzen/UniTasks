<template #app>
  <va-navbar color="primary">
    <template #left>
      <va-navbar-item class="logo">UniTasks</va-navbar-item>
    </template>
    <template #center>
      <va-navbar-item>All Tasks</va-navbar-item>
      <va-navbar-item>Tasks</va-navbar-item>
      <va-navbar-item>Recurring Tasks</va-navbar-item>
    </template>
    <template #right>
      <va-navbar-item>
        <va-button-dropdown icon="mdi-account_circle">
          <va-button >Settings</va-button> <br>
        <va-button >Logout</va-button>
        </va-button-dropdown>   
      </va-navbar-item>
    </template>
  </va-navbar>
</template>

<script>

export default {
  name: 'App',
  components: {
    
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
</style>
