package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Carjul/Noticias/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEstadosNotice(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("EstadoArticulo")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	var estadoNotice []bson.M

	if err = cursor.All(context.Background(), &estadoNotice); err != nil {
		panic(err)
	}
	responseJSON, err := json.Marshal(estadoNotice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar los estadoNotice como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

func GetEstadoNotice(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("EstadoArticulo")

	vars := mux.Vars(r)
	id := vars["id"]
	// Convertir el ID de string a ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Realizar la búsqueda del artículo por el filtro
	var EstNotice bson.M
	err = coll.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&EstNotice)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Rol no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error al buscar el rol: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(EstNotice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar el rol como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

func CreateEstadoNotice(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("EstadoArticulo")

	var EstNotice bson.M
	err := json.NewDecoder(r.Body).Decode(&EstNotice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer el rol: %s", err.Error()), http.StatusBadRequest)
		return
	}

	result, err := coll.InsertOne(context.Background(), EstNotice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al insertar el rol: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	EstNotice["ID"] = result.InsertedID

	responseJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar el rol como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func DeleteEstadoNotice(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("EstadoArticulo")

	vars := mux.Vars(r)
	id := vars["id"]
	// Convertir el ID de string a ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := coll.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar el rol: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Artículo no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"deleted": true}`))
}
