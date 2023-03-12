<template>
    <va-navbar-item>
        <va-dropdown v-if="initial() != null">
            <template #anchor>
                <va-button>
                    <va-avatar size="small" color="secondary" style="font-size: 100%;">{{ initial() }}</va-avatar>
                </va-button> 
            </template>
            <va-dropdown-content>
                <va-button >Settings</va-button> <br>
                <va-button >Logout</va-button>
            </va-dropdown-content>
        </va-dropdown>  
        <va-dropdown v-else>
            <template #anchor>
                <va-button>
                    <va-avatar size="small" color="secondary">
                        <va-icon name="mdi-account_circle"></va-icon>
                    </va-avatar>
                </va-button> 
            </template>
            <va-dropdown-content>
                <va-button >Login</va-button> <br>
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