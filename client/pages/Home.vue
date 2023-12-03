<template>
  <div>
    <h1>Welcome to the Home Page</h1>
    <div class="" v-for="i in PubArt" :key="i._id">

      <div class="card mb-3" style="max-width: 540px;">
        <div class="row g-0">
          <!-- <div class="col-md-4">
            <img :src="i.Imagen" class="img-fluid rounded-start" alt="...">
          </div> -->
          <div class="col-md-8">
            <div class="card-body">
              <h5 class="card-title">{{i.Titulo}}</h5>
              <p class="card-text">{{i.Cuerpo}}</p>
              <p class="card-text"><small class="text-muted">{{i.Fecha}}</small></p>
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup >
import { ref,onMounted } from 'vue'
import { useDataStore } from '../stores/data'

const dataStore = useDataStore()

onMounted(() => {
  fetch('http://localhost:8000/articulos')
  .then(res => res.json())
  .then(data => {
    if(data.length > 0){
      dataStore.setNoticias(data)
    let filtro = data.filter((articulo) => {return articulo.Estado.Nombre == "Publicado"})
      PubArt.value = filtro
    }
})


})

const PubArt = ref([])

watch(PubArt, (item) => {
  console.log(item)
  console.log(dataStore.noticias)
  
}) 

</script>

<style scoped>
/* Add your component styles here */
</style>
