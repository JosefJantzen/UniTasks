<template>
    Login
    <va-button @click="login"> Login </va-button>
</template>

<script>
import api from '../api/apiClient'

export default {
    name: 'Login',
    methods: {
        async login() {
            await api.post('/signIn', {
                "eMail": "admin@admin.com",
                "pwd": "secret"
            }).then(() => {
                if (this.$route.query.redirect && this.$route.query.redirect.indexOf('/') === 0) {
                    this.$router.push(this.$route.query.redirect)
                } else {
                    this.$router.push('/')
                }
            })
            
        }
    }
}

</script>