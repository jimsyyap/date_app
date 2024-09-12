<template>
    <div class="app">
        <h1>user data from postgres</h1>
        <ul v-if="users.length">
            <li v-for="user in users" :key="user.id">
                {{ user.name }} (ID: {{ user.id }})
            </li>
        </ul>
        <p v-else-if="loading">Loading...</p>
        <p v-else>No users found.</p>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
    name: 'App',
    setup() {
        const users = ref([])
        const loading = ref(true)
        const fetchUsers = async() => {
            try {
                const response = await fetch('/api/data')
                const data = await response.json()
                users.value = data
            } catch (error) {
                console.error('Error fetching users:', error)
            } finally {
                loading.value = false
            }
        }

        onMounted(fetchUsers)
        return {
            users,
            loading
        }
    }
}
</script>

<style>
.app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}
</style>
