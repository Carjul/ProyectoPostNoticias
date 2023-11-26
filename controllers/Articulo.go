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

func GetArticulos(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("Articulo")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	var articulos []bson.M

	if err = cursor.All(context.Background(), &articulos); err != nil {
		panic(err)
	}
	responseJSON, err := json.Marshal(articulos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar los artículos como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

func GetOneArticulo(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("Articulo")

	vars := mux.Vars(r)
	id := vars["id"]
	// Convertir el ID de string a ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Realizar la búsqueda del artículo por el filtro
	var articulo bson.M
	err = coll.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&articulo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Artículo no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error al buscar el artículo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Codificar el artículo como JSON y enviar la respuesta
	responseJSON, err := json.Marshal(articulo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar el artículo como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func CreateArticulo(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("Articulo")

	var articulo map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&articulo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer el artículo: %s", err.Error()), http.StatusBadRequest)
		return
	}

	result, err := coll.InsertOne(context.Background(), articulo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al insertar el artículo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Obtener el ID del artículo insertado
	articulo["ID"] = result.InsertedID

	// Codificar el artículo como JSON y enviar la respuesta
	responseJSON, err := json.Marshal(articulo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar el artículo como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)

}

func UpdateArticulo(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("Articulo")

	vars := mux.Vars(r)
	id := vars["id"]
	// Convertir el ID de string a ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var articulo map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&articulo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer el artículo: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Eliminar el ID del artículo para evitar que se actualice
	delete(articulo, "ID")

	result, err := coll.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": articulo})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar el artículo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Artículo no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"updated": true}`))
}

func DeleteArticulo(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("Articulo")

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
		http.Error(w, fmt.Sprintf("Error al eliminar el artículo: %s", err.Error()), http.StatusInternalServerError)
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
