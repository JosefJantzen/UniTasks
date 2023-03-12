<template>
    <va-navbar-item>
        <va-dropdown 
            v-if="initial() != null"
            placement="bottom-end"
            :offset="[10, 0]"
        >
            <template #anchor>
                <va-button>
                    <va-avatar size="small" color="secondary" style="font-size: 100%;">{{ initial() }}</va-avatar>
                </va-button> 
            </template>
            <va-dropdown-content class="drop">
                <va-button class="drop-btn" preset="secondary" icon="mdi-settings">Settings</va-button> 
                <br>
                <br>
                <va-button class="drop-btn" preset="secondary" icon="mdi-logout">Logout</va-button>
            </va-dropdown-content>
        </va-dropdown>  
        <va-button v-else class="drop-btn" color="info" icon="mdi-login">Login</va-button> 
    </va-navbar-item>
</template>

<script>
export default {
    name: 'AccountMenu',
    methods: {
        initial() {
            let user = this.$store.getters['user/get']
            console.log(user)
            if (user != null && user.eMail != null) {
                return user.eMail.charAt(0).toUpperCase()
            }
            return null
        }
    }
}
</script>

<style>
.drop {
    padding-top: 1em;
    padding-bottom: 1em;
}

.drop-btn {
    margin-left: 1rem;
    margin-right: 1rem;
}
</style>