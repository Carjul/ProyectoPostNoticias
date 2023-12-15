
import {useDataStore} from '../stores/data'
const x = useDataStore()

export default defineNuxtRouteMiddleware((to, from) => {
    if (x.isAuthtenticated === false) {
      return navigateTo('/')
    }
  })