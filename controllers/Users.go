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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	collUsers := db.Client.Database("Noticias").Collection("User")
	collRoles := db.Client.Database("Noticias").Collection("Rol")

	cursor, err := collUsers.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al buscar usuarios: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var users []bson.M

	if err = cursor.All(context.Background(), &users); err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener usuarios: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Recorrer cada usuario para obtener y asociar la información del rol
	for _, user := range users {
		rolID, ok := user["RolId"].(primitive.ObjectID)
		if !ok {
			http.Error(w, "ID de rol no válido", http.StatusInternalServerError)
			return
		}

		var rol bson.M
		err = collRoles.FindOne(context.TODO(), bson.M{"_id": rolID}).Decode(&rol)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error al buscar el rol: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// Asociar la información del rol al usuario
		user["Rol"] = rol
	}

	responseJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar los usuarios como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	collUsers := db.Client.Database("Noticias").Collection("User")
	collRoles := db.Client.Database("Noticias").Collection("Rol")

	vars := mux.Vars(r)
	id := vars["id"]

	// Convertir el ID de string a ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Realizar la búsqueda del usuario por el filtro
	var user bson.M
	err = collUsers.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error al buscar el usuario: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Obtener el RolID del usuario
	rolID, ok := user["RolId"].(primitive.ObjectID)
	if !ok {
		http.Error(w, "ID de rol no válido", http.StatusInternalServerError)
		return
	}

	// Realizar la búsqueda del rol por su ID
	var rol bson.M
	err = collRoles.FindOne(context.TODO(), bson.M{"_id": rolID}).Decode(&rol)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Rol no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error al buscar el rol: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Agregar el detalle del rol al usuario
	user["Rol"] = rol

	responseJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar el usuario como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("User")
	Rol := db.Client.Database("Noticias").Collection("Rol")

	var rol bson.M
	err1 := Rol.FindOne(context.TODO(), bson.M{"Nombre": "Visitante"}).Decode(&rol)
	if err1 != nil {
		if err1 == mongo.ErrNoDocuments {
			http.Error(w, "Rol no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error al buscar el Rol: %s", err1.Error()), http.StatusInternalServerError)
		return
	}
	objID, err2 := primitive.ObjectIDFromHex(rol["_id"].(primitive.ObjectID).Hex())
	if err2 != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var user bson.M
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer el body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	user["RolId"] = objID

	result, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al insertar el usuario: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al codificar el ID del usuario como JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("User")

	vars := mux.Vars(r)
	id := vars["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Supongamos que se espera un JSON con los campos a actualizar en la solicitud
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, fmt.Sprintf("Error al decodificar los datos de actualización: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Eliminar el campo 'rol_id' del mapa de actualización para evitar cambios en el rol
	delete(updateData, "RolId")

	// Crear un filtro para encontrar el usuario a actualizar por su ID
	filter := bson.M{"_id": objID}

	// Crear una actualización con los campos recibidos
	update := bson.M{"$set": updateData}

	// Realizar la actualización del usuario
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar el usuario: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Verificar si no se encontró el usuario
	if result.MatchedCount == 0 {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Usuario actualizado correctamente")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	coll := db.Client.Database("Noticias").Collection("User")

	vars := mux.Vars(r)
	id := vars["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := coll.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar el usuario: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Usuario eliminado correctamente")
}

func UpdateRolUser(w http.ResponseWriter, r *http.Request) {
	type data struct {
		ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		NombreRol string             `json:"Nombre_rol"`
	}

	var datos data
	err := json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer el body: %s", err.Error()), http.StatusBadRequest)
		return

	}
	colrol := db.Client.Database("Noticias").Collection("Rol")
	coluser := db.Client.Database("Noticias").Collection("User")

	var rol bson.M

	err1 := colrol.FindOne(context.Background(), bson.M{"Nombre": datos.NombreRol}).Decode(&rol)
	if err1 != nil {
		if err1 == mongo.ErrNoDocuments {
			http.Error(w, "Rol no encontrado", http.StatusNotFound)
			return
		}
	}
	objID, err2 := primitive.ObjectIDFromHex(rol["_id"].(primitive.ObjectID).Hex())
	if err2 != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := coluser.UpdateOne(context.TODO(), bson.M{"_id": datos.ID}, bson.M{"$set": bson.M{"RolId": objID}})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar el usuario: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "rol de Usuario actualizado correctamente")

}
