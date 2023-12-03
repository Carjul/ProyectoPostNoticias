<template>
  <div>

    <div>
      <section>
        <section class="sticky">

          <div class="max-w-lg px-4 sm:pt-24 pt-12 sm:pb-8 mx-auto text-left md:max-w-none md:text-center">



            <h1
              class="font-extrabold leading-10 tracking-tight text-left text-[#201515] text-center sm:leading-none text-5xl sm:text-9xl">
              <span class="inline md:block">Building Good </span>
              <span class="relative mt-2 bg-clip-text text-[#201515] md:inline-block">Software.</span>
            </h1>
          </div>


          <div class="max-w-lg px-4 pb-24 mx-auto text-left md:max-w-none md:text-center">
            <div class='text-center py-4 space-x-4' v-if="mostrar">

              <button @click="toggle" v-if="mostrarbtn"
                class="backdrop-blur-sm transition duration-500 ease-in-out bg-[#FF4F01] border border-[#E2E8F0] translate-y-1 text-white hover:bg-neutral-200 text-lg font-semibold py-3 px-6 rounded-3xl inline-flex items-center">
                <span> LogIn</span>
              </button>
              <NuxtLink to="/home">
                <button
                  class="backdrop-blur-sm transition duration-500 ease-in-out bg-white border border-[#E2E8F0] translate-y-1 text-[#16161d] hover:bg-neutral-200 text-lg font-semibold py-3 px-6 rounded-3xl inline-flex items-center">
                  <span> Ver</span>
                </button>
              </NuxtLink>
            </div>
            <section class="flex justify-center items-center" v-else>
              <div class="max-w-md w-full bg-white rounded p-6 space-y-4">
                <div class="flex flex-row justify-between">
                  <p class="text-gray-600">Sign In</p>
                  <button @click="toggle"
                    class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-500 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm">
                    x</button>

                </div>

                <div>
                  <input v-model="sendUser.correo"
                    class="w-full p-4 text-sm bg-gray-50 focus:outline-none border border-gray-200 rounded text-gray-600"
                    type="text" placeholder="Email">
                </div>
                <div>
                  <input v-model="sendUser.password"
                    class="w-full p-4 text-sm bg-gray-50 focus:outline-none border border-gray-200 rounded text-gray-600"
                    type="text" placeholder="Password">
                </div>
                <div>
                  <button @click="login"
                    class="w-full py-4 bg-blue-600 hover:bg-blue-700 rounded text-sm font-bold text-gray-50 transition duration-200">Sign
                    In
                  </button>
                </div>
                <div class="flex items-center justify-end">

                  <div>
                    <a class="text-sm text-blue-600 hover:underline" href="#">Forgot password?</a>
                  </div>
                </div>
              </div>
            </section>

          </div>


        </section>
      </section>


      <div class="text-left">

        <div class='sm:px-28'>
          <section class="relative flex items-center w-full">
            <div class="relative items-center w-full px-5 mx-auto md:px-12 lg:px-16 max-w-7xl">
              <div class="relative flex-col items-start m-auto align-middle">
                <div class="grid grid-cols-1 gap-6 lg:grid-cols-2 lg:gap-24">
                  <div class="relative items-center gap-12 m-auto lg:inline-flex md:order-first">
                    <div class="max-w-xl text-center lg:text-left">
                      <div>

                        <p class="text-3xl font-semibold tracking-tight text-[#201515] sm:text-5xl">
                          Space Management Software
                        </p>
                        <p class="max-w-xl mt-4 text-base tracking-tight text-gray-600">
                          Use this paragraph to share information about your company or products. Make it engaging and
                          interesting, and showcase your brand's personality. Thanks for visiting our website!
                        </p>
                      </div>
                      <div class="flex justify-center gap-3 mt-10 lg:justify-start">
                        <a class="inline-flex items-center justify-center text-sm font-semibold text-black duration-200 hover:text-blue-500 focus:outline-none focus-visible:outline-gray-600"
                          href="#">
                          <span> Read more &nbsp; → </span>
                        </a>
                      </div>
                    </div>
                  </div>
                  <div class="order-first block w-full mt-12 aspect-square lg:mt-0">
                    <img class="object-cover rounded-3xl object-center w-full mx-auto bg-gray-300 lg:ml-auto " alt="hero"
                      src="https://i.pinimg.com/originals/2e/2b/21/2e2b21aeed393403d4620367f9e093f9.gif" />
                  </div>
                </div>
              </div>
            </div>
          </section>

        </div>

        <div class="mt-32" />

        <section>

        </section>

      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref,onMounted } from 'vue'
import { useDataStore } from '../stores/data'

const dataStore = useDataStore()

 onMounted(() => {
    if (dataStore.user.hasOwnProperty('_id')) {
    mostrarbtn.value = false
  }else{
    mostrarbtn.value = true
  }
  })

useHead({
  title: 'App Notices',
  meta: [
    { name: 'description', content: 'My amazing site.' }
  ],
  bodyAttrs: {
    class: 'test'
  }
})

const sendUser = ref({
  correo: "",
  password: ""
})
const mostrarbtn = ref(true)
const mostrar = ref(true)

const login = async () => {

  const res = await fetch('http://localhost:8000/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(sendUser.value)
  })
  const data = await res.json()

  if (data != "Usuario o contraseña incorrectos") {
    dataStore.setUser(data)
  } else {
    alert(data)
  }

  if (dataStore.user.hasOwnProperty('_id')) {
    mostrarbtn.value = false
    await navigateTo('/home')
  }
}



const toggle = () => {
  if (mostrar.value == true) {
    mostrar.value = false
  } else {
    mostrar.value = true
  }
}

</script>

<style scoped></style>
