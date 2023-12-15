import { defineStore } from 'pinia'

export const useDataStore = defineStore('data', {
  state: () => {
    return { 
        user: {},
        noticias: [],
        isAuthtenticated: false,
        }
  },
  
  actions: {
    setUser(user) {
      this.user = user
    },
    setDeleteUser() {
        this.user = {}
    },
    setNoticias(noticias) { 
      this.noticias=noticias;
    },
    setDeleteNoticias() {
      this.noticias = []
    },
    setAuthtenticated(params) {
      this.isAuthtenticated = params
    },
  },
})