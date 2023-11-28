<template>
  <div>
    <h1 class="text-3xl font-bold underline">Welcome to the Main View!</h1>
    <!-- Add your content here -->
  </div>
</template>

<script setup lang="ts">
useHead({
  title: 'App Notices',
  meta: [
    { name: 'description', content: 'My amazing site.' }
  ],
  bodyAttrs: {
    class: 'test'
  }
})
</script>

<style scoped>
/* Add your component styles here */
</style>
