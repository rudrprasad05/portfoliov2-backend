package main

import (
	"log"
	"net/http"

	"portfolio_backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rudrprasad05/go-logs/logs"
	"github.com/rudrprasad05/go-sql/connect"
)

type Message struct {
	Data string `json:"data"`
}

func main() {
	
	router := mux.NewRouter()
	config := connect.Config{
		Username: "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "portfolio",
	}

	logger, err := logs.NewLogger()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logger.Close()

	// Initialize the database
	db, err := connect.InitDB(&config)
	routes := &routes.Routes{DB: db, LOG: logger}
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
		logger.Error("failed to create db")
	}
	defer db.Close()

	// routes
	router.HandleFunc("/404", routes.Handle404)
	router.HandleFunc("/", routes.GetHome).Methods("GET")
	router.HandleFunc("/auth/register", routes.PostRegisterUser).Methods("POST")
	router.HandleFunc("/auth/login", routes.PostLoginUser).Methods("POST")

	protected := router.PathPrefix("/protected").Subrouter()
	protected.Use(routes.AuthMiddleware)

	protected.HandleFunc("", routes.GetProtectedAuth).Methods("GET")

	// redirect handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
	})
	// middlewaree
	corsRouter := routes.CorsMiddleware(router)
	loggedHandler := logs.LoggingMiddleware(logger, corsRouter)

	log.Fatal(http.ListenAndServe(":8080", loggedHandler))
}