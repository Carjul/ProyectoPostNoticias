package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Carjul/Noticias/controllers"
	"github.com/Carjul/Noticias/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	/* Cargar env */
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	/* conectar db */
	db.ConexionDB()

	app := mux.NewRouter()

	/* Rutas */
	app.HandleFunc("/", controllers.Init).Methods("GET")
	//Articulos CRUD
	app.HandleFunc("/articulos", controllers.GetArticulos).Methods("GET")
	app.HandleFunc("/articulo/{id}", controllers.GetOneArticulo).Methods("GET")
	app.HandleFunc("/articulo", controllers.CreateArticulo).Methods("POST")
	app.HandleFunc("/articuloEst", controllers.UpdateEstArticulo).Methods("POST")
	app.HandleFunc("/articulo/{id}", controllers.UpdateArticulo).Methods("PUT")
	app.HandleFunc("/articulo/{id}", controllers.DeleteArticulo).Methods("DELETE")
	//EstadoArticulos CRUD
	app.HandleFunc("/estados", controllers.GetEstadosNotice).Methods("GET")
	app.HandleFunc("/estado/{id}", controllers.GetEstadoNotice).Methods("GET")
	app.HandleFunc("/estado", controllers.CreateEstadoNotice).Methods("POST")
	app.HandleFunc("/estado/{id}", controllers.DeleteEstadoNotice).Methods("DELETE")
	//Roles CRUD
	app.HandleFunc("/roles", controllers.GetRoles).Methods("GET")
	app.HandleFunc("/rol/{id}", controllers.GetOneRol).Methods("GET")
	app.HandleFunc("/rol", controllers.CreateRol).Methods("POST")
	app.HandleFunc("/rol/{id}", controllers.DeleteRol).Methods("DELETE")
	//Users CRUD
	app.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	app.HandleFunc("/user/{id}", controllers.GetOneUser).Methods("GET")
	app.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	app.HandleFunc("/userRol", controllers.UpdateRolUser).Methods("POST")
	app.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	app.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	/* Cors */
	Cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
	})

	/* Envolver el enrutador con el middleware CORS */
	router := Cors.Handler(app)

	/* SET PORT */
	log.Println("Server run on port " + port + "")
	http.ListenAndServe(":"+port, router)
}
