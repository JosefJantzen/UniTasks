<template>
    <h1 class="heading">Settings</h1>
    <va-card class="card">
        <va-card-title>
            EMail Settings
        </va-card-title>
        <va-card-content>
            <va-form
                tag="form"
                @submit="submitMail"
            >
                <va-input
                class="inputs"
                label="Old EMail"
                type="email"
                v-model="oldEmail"
                /><br>
                <va-input
                class="inputs"
                label="New EMail"
                type="email"
                v-model="newEmail1"
                /><br>
                <va-input
                class="inputs"
                label="Repeat new EMail"
                type="email"
                v-model="newEmail2"
                :rules="[(newEmail1 == newEmail2)]"
                /><br>
                <va-button
                    type="submit"
                >Change email</va-button>
            </va-form>
        </va-card-content>
    </va-card>
    <va-card class="card" style="margin-top: 20px;">
        <va-card-title>
            Password Settings
        </va-card-title>
        <va-card-content>
            <va-form
                tag="form"
                @submit="changeMail"
            >
                <va-input
                class="inputs"
                label="Old Password"
                type="password"
                v-model="oldPwd"
                /><br>
                <va-input
                class="inputs"
                label="New Password"
                type="password"
                v-model="newPwd1"
                /><br>
                <va-input
                class="inputs"
                label="Repeat new Password"
                type="password"
                v-model="newPwd2"
                :rules="[(newPwd1 == newPwd2)]"
                /><br>
                <va-button
                    type="submit"
                >Change password</va-button>
            </va-form>
        </va-card-content>
    </va-card>
</template>

<script>
import { mapActions } from 'vuex'
import { useToast } from 'vuestic-ui/web-components'

export default {
    name: 'Settings',
    methods: {
        ...mapActions('user', ['changeMail']),
        async submitMail() {
            if (this.oldEmail != this.$store.getters['user/get'].eMail) {
                useToast().init({
                    title: "Email change failed",
                    message: "Old Email incorrect",
                    color: 'danger',
                    position: 'bottom-right',
                    duration: 3000

                })
                return
            }
            if (this.newEmail1 != this.newEmail2) {
                useToast().init({
                    title: "Email change failed",
                    message: "New Emails are different",
                    color: 'danger',
                    position: 'bottom-right',
                    duration: 3000

                })
                return
            }
            await this.changeMail(this.newEmail1)
            useToast().init({
                title: "Email changed",
                message: "Your email was succesfully changed",
                color: 'success',
                position: 'bottom-right',
                duration: 3000

            })
        }
    },
    data () {
        return {
            oldEmail: "",
            newEmail1: "",
            newEmail2: "",
            oldPwd: "",
            newPwd1: "",
            newPwd2: ""
        }
    }
}
</script>

<style>
.heading {
    margin-top: 110px;
    font-size:xx-large;
    margin-bottom: 1rem;
    position: relative;
    z-index: 105;
}

.card {
    z-index: 105;
    width: 50%;
    position: center;
    margin: 0 auto;
}

.inputs {
    margin-bottom: 1em;
}

</style>