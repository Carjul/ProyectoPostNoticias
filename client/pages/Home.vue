<template>
  <div class="flex flex-col w-auto h-auto" id="home">
    <Nav></Nav>
    <div class="flex flex-row justify-center">
      <Card v-for="i in PubArt" :key="i._id" :Img="i.Img" :Titulo="i.Titulo" :Cuerpo="i.Cuerpo"></Card>
    </div>
  </div>
</template>

<script setup>
import { ref,onMounted } from 'vue'
import { useDataStore } from '../stores/data'
useHead({
  title: 'Home',
  meta: [
    { name: 'Home', content: 'Home' }
  ],
  link: [
    { rel: 'shortcut icon', href: 'https://i.pinimg.com/originals/2e/2b/21/2e2b21aeed393403d4620367f9e093f9.gif' }
  ]
  
})



const Env =useRuntimeConfig()

const dataStore = useDataStore()

onMounted(() => {
  fetch(`${Env.public.apiUrl}/articulos`)
  .then(res => res.json())
  .then(data => {
    if(data.length > 0){
      dataStore.setNoticias(data)
    let filtro = data.filter((articulo) => {return articulo.Estado.Nombre === "Publicado"})
      PubArt.value = filtro
    }
})


})

const PubArt = ref([])

/* watch(PubArt, (item) => {
  console.log(item)
  console.log(dataStore.noticias)
  
}) 
 */
</script>

<style scoped>

</style>
