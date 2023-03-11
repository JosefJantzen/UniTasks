<template>
    <va-button @click="foo"> Button </va-button>
</template>

<script>
export default {
    name: 'TaskDashboard',
    methods: {
        foo() {
            console.log(api.post('/signin'))
        }
    }
}

import api from '../api/apiClient'

</script>