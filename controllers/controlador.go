package controllers

import (
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bienvenido a la API de Noticias"))

}
