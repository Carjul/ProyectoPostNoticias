import { defineStore } from 'pinia'

export const useDataStore = defineStore('data', {
  state: () => {
    return { 
        user: {},
        noticias: [],
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
  },
})