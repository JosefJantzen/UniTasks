<template>
    <va-image class="logo" src="logo.png" :max-width=80 style="margin-top: 110px"/>
    <h1 class="heading">Sign up</h1>
    <va-form 
        class="sign-up-form"
        tag="form"
        @submit.prevent="submit"
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
            :rules="[(v) => v.length >= 8 || `At least 8 characters`]"
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
        <va-button type="submit"> Sign up </va-button>
    </va-form>
    <br>
    <va-button @click="$router.replace('/login')" preset="plain">Login instead</va-button>
</template>

<script>
import { useToast } from 'vuestic-ui/web-components'
import { mapActions } from 'vuex'

export default {
    name: 'SignUp',
    methods: {
        ...mapActions('user', ['signUp']),
        async submit() {
            if (!/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(this.$data.email)) {
                toast("Please enter valid email")
                return
            }
            if (this.$data.pwd.length < 8) {
                toast("Password needs at least 8 chracter")
                return
            }
            try {
                await this.signUp({
                    "eMail": this.$data.email,
                    "pwd": this.$data.pwd,
                })
                if (this.$route.query.redirect && this.$route.query.redirect.indexOf('/') === 0) {
                    this.$router.replace(this.$route.query.redirect)
                } else {
                    this.$router.replace('/')
                }
            } catch (e) {
                toast("Email already used")
            }
        }
    },    
    data: () => ({
        isPasswordVisible: false,
        email: "",
        pwd: "",
    }),
}

function toast(msg) {
    useToast().init({
        title: "Sign up failed",
        message: msg,
        color: 'danger',
        position: 'bottom-right',
        duration: 3000

    })
}

</script>