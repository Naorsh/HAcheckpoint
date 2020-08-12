package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"checkpointapp/internal/pkg/db"
	"checkpointapp/internal/pkg/models"

	"github.com/gorilla/mux"
)

var httpServer *http.Server

const authHeader = "Cookie"
const authSecret = "CHECKPOINTID=let-me-pass"

// Middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authSecret != r.Header.Get(authHeader) {
			log.Println("Unauthorized attempt blocked")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Start starts the web server
func Start(port string) error {

	r := mux.NewRouter()
	healthRouter := r.PathPrefix("/").Subrouter()
	healthRouter.HandleFunc("/live", HealthHandler)

	appsRouter := r.PathPrefix("/api").Subrouter()
	appsRouter.HandleFunc("/addApplication", createApplication)
	appsRouter.HandleFunc("/getApplications", getApplications)
	appsRouter.Use(authMiddleware)

	httpServer = &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("HTTP server up on %s", httpServer.Addr)
	return httpServer.ListenAndServe()
}

//  gets ALL applications from DB
func getApplications(w http.ResponseWriter, r *http.Request) {
	var apps []models.App
	apps, err1 := db.GetApplications()
	if err1 != nil {
		log.Println(err1)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apps)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

//  isnerting Application to DB
func createApplication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var app models.App
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	app = models.NewApp(app.Name, app.Key)

	//inpur validation
	if app.Name == "" || app.Key == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	//insert value
	err = db.InsertNewApp(app)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// HealthHandler handles the "/health" route
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("CHECKPOINTID")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "CHECKPOINTID",
			Value: "let-me-pass",
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	log.Println("Health called")
	log.Println("Request:", r)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}
