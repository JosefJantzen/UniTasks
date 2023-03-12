<template>
    <va-image class="logo" src="logo.png" :max-width=80 />
    <h1 class="heading">Sign in</h1>
    <div class="login-form">
        <va-input
            class="inputs"
            label="EMAIL"
            type="email"
            v-model="email"
        /><br>
        <va-input
            class="inputs"
            v-model="pwd"
            label="PASSWORD"
            :type="isPasswordVisible ? 'text' : 'password'"
        >
            <template #appendInner>
                <va-icon
                    :name="isPasswordVisible ? 'visibility_off' : 'visibility'"
                    size="small"
                    color="--va-primary"
                    @click="isPasswordVisible = !isPasswordVisible"
                />
            </template>
        </va-input>
        <br>
        <va-button @click="login"> Login </va-button>
    </div>
    
</template>

<script>
import api from '../api/apiClient'

export default {
    name: 'Login',
    methods: {
        async login() {
            await api.post('/signIn', {
                "eMail": this.$data.email,
                "pwd": this.$data.pwd,
            }).then(() => {
                
                if (this.$route.query.redirect && this.$route.query.redirect.indexOf('/') === 0) {
                    this.$router.push(this.$route.query.redirect)
                } else {
                    this.$router.push('/')
                }
            }).catch(() => {
                console.log("Failed login")
            })
            
        }
    },
    data: () => ({
        isPasswordVisible: false,
        email: "",
        pwd: "",
    }),
}

</script>


<style>
.heading {
    font-size:xx-large;
    margin-bottom: 1rem;
}

.login-form {
    margin-top: 2em;
}

.logo {
    margin-left: auto;
    margin-right: auto;
}

.inputs{
    margin-bottom: 1em;
}
</style>