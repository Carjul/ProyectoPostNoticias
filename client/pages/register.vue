<template>
  <div>
    <h1>Registration</h1>
    <form @submit="submitForm">
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="username" required>
      
      <label for="email">Email:</label>
      <input type="email" id="email" v-model="email" required>
      
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" required>
      
      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      email: '',
      password: ''
    }
  },
  methods: {
    submitForm() {
      // Perform registration logic here
      // You can use this.username, this.email, this.password to access the form values
    }
  }
}
</script>

<style scoped>
/* Add your custom styles here */
</style>
