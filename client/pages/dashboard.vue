<template>
  <div class="flex flex-col">
    <Nav></Nav>
    <div id="admin" v-if="rol.Nombre === 'Admistrador'"> hola soy admin</div>
    <div id="redactor"  v-if="rol.Nombre === 'Redactor'">hola soy Redactor</div>
    <div id="editor"  v-if="rol.Nombre === 'Editor'">hola soy editor</div>
  </div>
</template>

<script setup>
import {useDataStore} from '../stores/data'
import { ref, onMounted } from 'vue'

const Env =useRuntimeConfig()
const dataStore = useDataStore()
var rol = ref({Nombre:''})

definePageMeta({middleware:'auth'})

onMounted(() => {
  if (dataStore.user.hasOwnProperty('_id') ) {
   fetch(`${Env.public.apiUrl}/user/${dataStore.user._id}`)
    .then(res => res.json())
    .then(data => {
      if(data.Rol.Nombre){
        rol.value.Nombre = data.Rol.Nombre
      }
    })
}
})
 
useHead({
  title: 'Dashboard',
  meta: [
    { name: 'description', content: 'Dashboard' }
  ],
  link: [
    { rel: 'shortcut icon', href: 'https://i.pinimg.com/originals/2e/2b/21/2e2b21aeed393403d4620367f9e093f9.gif' }
  ]
})




</script>

<style scoped>
/* Add your component styles here */
</style>
