<template>
    <va-navbar-item>
        <va-dropdown 
            v-if="initial() != null"
            :placement="this.isMobile ? 'bottom' : 'bottom-end'"
            :offset="[10, 0]"
            trigger="hover"
            hover-out-timeout="500"
        >
            <template #anchor>
                <va-button>
                    <va-avatar :size=" this.isMobile ? '25px' : 'small'" color="info" style="font-size: 100%;">{{ initial() }}</va-avatar>
                </va-button> 
            </template>
            <va-dropdown-content :class="this.isMobile ? 'drop-mobile' : 'drop'">
                <va-button class="drop-btn" preset="secondary" icon="mdi-settings"
                    @click="$router.push('/settings')"
                    >Settings</va-button> 
                <br>
                <va-button class="drop-btn" preset="secondary" icon="mdi-logout"
                    @click="logout">Logout</va-button>
            </va-dropdown-content>
        </va-dropdown>  
        <va-button  v-else class="drop-btn login" color="info" icon="mdi-login"
            @click="$router.push('/login')">Login</va-button> 
    </va-navbar-item>
</template>

<script>
import { mapActions } from 'vuex'
import help from '../../help/help'

export default {
    name: 'AccountMenu',
    methods: {
        ...mapActions('user', ['logout']),
        initial() {
            let user = this.$store.getters['user/get']
            if (user != null && user.eMail != null) {
                return user.eMail.charAt(0).toUpperCase()
            }
            return null
        }
    },
    data () {
        return {
            isMobile: help.isMobile()
        }
    }
}
</script>

<style>
.drop {
    padding-top: 1rem;
    padding-bottom: 1rem;
    z-index: 201;
}

.drop-mobile {
    padding-top: 1rem;
    padding-bottom: 1rem;
    z-index: 201;
}

.drop-btn {
    margin-left: 1rem;
    margin-right: 1rem;
}

.login {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
}
</style>