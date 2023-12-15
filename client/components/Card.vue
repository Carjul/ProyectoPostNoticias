<template>
    <div class="flex flex-col items-center justify-center" id="card">
        <div class="max-w-sm rounded overflow-hidden shadow-lg">
          <img class="w-full" :src="Img" alt="Imagen">
          <div class="px-6 py-4">
            <div class="font-bold text-xl mb-2">{{Titulo}}</div>
            <p class="text-gray-700 text-base">{{Cuerpo}}</p>
          </div>
        </div>
      </div>
</template>

<script setup>
const props = defineProps(['Img','Titulo','Cuerpo'])

</script>

<style>
#card{
    width: 100%;
    height:auto;
    margin-bottom: 10px;
}
</style>
