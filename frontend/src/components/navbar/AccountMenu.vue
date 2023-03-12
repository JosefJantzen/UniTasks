<template>
    <va-navbar-item>
        <va-dropdown>
            <template #anchor>
                <va-button @click="foo">
                    <va-avatar >{{ initial() }}</va-avatar>
                </va-button>
                
            </template>
            <va-dropdown-content>
                <va-button >Settings</va-button> <br>
                <va-button >Logout</va-button>
            </va-dropdown-content>
        </va-dropdown>   
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