<template>
    <va-image class="logo" src="logo.png" :max-width=80 style="margin-top: 110px"/>
    <h1 class="heading">Sign in</h1>
    <va-form 
        class="login-form"
        tag="form"
        @submit.prevent="login"
    >
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
        <va-button type="submit"> Login </va-button>
    </va-form>
    <br>
    <va-button @click="$router.replace('/signUp')" preset="plain">Create account</va-button>
</template>

<script>
import { useToast } from 'vuestic-ui/web-components'
import { mapActions } from 'vuex'

export default {
    name: 'Login',
    methods: {
        ...mapActions('user', ['signIn']),
        ...mapActions('user', ['refresh']),
        async login() {
            let intervalId = setInterval(() => this.refresh(), 20000)
            try {
                await this.signIn({
                    eMail: this.$data.email,
                    pwd: this.$data.pwd,
                    intervalId: intervalId
                })
                if (this.$route.query.redirect && this.$route.query.redirect.indexOf('/') === 0) {
                    this.$router.replace(this.$route.query.redirect)
                } else {
                    this.$router.replace('/')
                }
                  
            } catch (e) {
                clearInterval(intervalId)
                useToast().init({
                    title: "Login failed",
                    message: "Email or password wrong",
                    color: 'danger',
                    position: 'bottom-right',
                    duration: 3000

                })
            }
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
    position: relative;
    z-index: 105;
}

.login-form {
    margin-top: 2em;
    z-index: 105;
}

.logo {
    margin-left: auto;
    margin-right: auto;
    z-index: 105;
}

.inputs{
    margin-bottom: 1em;
    z-index: 105;
}
</style>