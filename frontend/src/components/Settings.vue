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
                :rules="[(newEmail1 == newEmail2)  || `Don't match`]"
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
                @submit="submitPwd"
            >
                <va-input
                class="inputs"
                label="New Password"
                v-model="newPwd1"
                :rules="[((v) => v.length >= 8 || `At least 8 characters`)]"
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
                </va-input><br>
                <va-input
                class="inputs"
                label="Repeat new Password"
                v-model="newPwd2"
                :rules="[((v) => v.length >= 8 || `At least 8 characters`), (newPwd1 == newPwd2) || `Don't match`]"
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
                </va-input><br>
                <va-button
                    type="submit"
                >Change password</va-button>
            </va-form>
        </va-card-content>
    </va-card>
    <va-card class="card" style="margin-top: 20px;">
        <va-card-title>
            Delete user account
        </va-card-title>
        <va-card-content>
            <va-form>
                <va-input
                    class="inputs"
                    label="EMAIL"
                    type="email"
                    v-model="delMail"
                /><br>
                <va-input
                    class="inputs"
                    v-model="delPwd"
                    label="PASSWORD"
                    :type="delIsPasswordVisible ? 'text' : 'password'"
                >
                    <template #appendInner>
                        <va-icon
                            :name="delIsPasswordVisible ? 'visibility_off' : 'visibility'"
                            size="small"
                            color="--va-primary"
                            @click="delIsPasswordVisible = !delIsPasswordVisible"
                        />
                    </template>
                </va-input><br>
                <va-button
                    color="danger"
                    @click="this.delete()"
                >Delete User</va-button>
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
        ...mapActions('user', ['changePwd']),
        ...mapActions('user', ['deleteUser']),
        async submitMail () {
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
        },
        async submitPwd () {
            if (this.newPwd1.length < 8) {
                return
            }
            if (this.newPwd1 != this.newPwd2) {
                useToast().init({
                    title: "Password change failed",
                    message: "New passowrds are different",
                    color: 'danger',
                    position: 'bottom-right',
                    duration: 3000

                })
                return
            }
            await this.changePwd(this.newPwd1)
            useToast().init({
                title: "Password changed",
                message: "Your password was succesfully changed",
                color: 'success',
                position: 'bottom-right',
                duration: 3000

            })
        },
        async delete () {
            this.$vaModal.init({
                title: 'Warning',
                message: 'Are you sure you want to delete this task?',
                okText: 'Yes',
                cancelText: 'No',
                blur: true,
                onOk: async () => {
                    try {
                        await this.deleteUser({
                            eMail: this.delMail,
                            pwd: this.delPwd
                        })
                    }
                    catch (e) {
                        useToast().init({
                            title: "Delete failed",
                            message: "Your credentials are wrong",
                            color: 'danger',
                            position: 'bottom-right',
                            duration: 3000

                        })
                    }
                },
            })
        }
    },
    data () {
        return {
            oldEmail: "",
            newEmail1: "",
            newEmail2: "",
            newPwd1: "",
            newPwd2: "",
            isPasswordVisible: false,
            delMail: "",
            delPwd: "",
            delIsPasswordVisible: false
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