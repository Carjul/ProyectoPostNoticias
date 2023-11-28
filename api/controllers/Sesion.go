package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Carjul/Noticias/db"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
)

var store = sessions.NewCookieStore([]byte("secret"))

func ProcessLogin(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}
	session, _ := store.Get(r, "session-name")

	var user User
	err2 := json.NewDecoder(r.Body).Decode(&user)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	coll := db.Client.Database("Noticias").Collection("User")

	var result bson.M
	errx := coll.FindOne(context.Background(), bson.M{"Correo": user.Correo, "Password": user.Password}).Decode(&result)
	if errx != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Usuario o contraseña incorrectos")
		return
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error al codificar los roles como JSON", http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	session.Values["authenticated"] = false
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Sesion cerrada")

}
