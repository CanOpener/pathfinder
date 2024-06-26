package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const defaultSearchSpaceGeneratorId = "ssgen"
const defaultGeneratorJobManagerId = "mutex"
const defaultPersistenceManagerId = "memory"

var mainSearchSpaceGenerator searchSpaceGenerator
var mainGenerationJobManager generationJobManager
var mainPersistenceManager persistenceManager

func main() {
	initializeSingletons()
	router := mux.NewRouter()

	router.Use(loggerMiddleware)

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router.HandleFunc("/generation_jobs", createGenerationJob).Methods("POST")
	router.HandleFunc("/generation_jobs/{id}", showGenerationJob).Methods("GET")

	router.HandleFunc("/search_spaces", indexSearchSpaces).Methods("GET")
	router.HandleFunc("/search_spaces", createSearchSpace).Methods("POST")
	router.HandleFunc("/search_spaces/{id}", showSearchSpace).Methods("GET")
	router.HandleFunc("/search_spaces/{id}", deleteSearchSpace).Methods("DELETE")

	log.Println("Server starting on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func initializeSingletons() {
	generator, err := newSearchSpaceGenerator(defaultSearchSpaceGeneratorId)
	if err != nil {
		log.Fatalf("failed to initialize main search space generator with id '%s' error: %v",
			defaultSearchSpaceGeneratorId, err)
	}
	mainSearchSpaceGenerator = generator

	jobManager, err := newGenerationJobManager(defaultGeneratorJobManagerId, mainSearchSpaceGenerator)
	if err != nil {
		log.Fatalf("failed to initialize main generation job manager with id '%s' error: %v",
			defaultGeneratorJobManagerId, err)
	}
	mainGenerationJobManager = jobManager

	persistenceManager, err := newPersistenceManager(defaultPersistenceManagerId)
	if err != nil {
		log.Fatalf("failed to initialize main persistence manager with id '%s' error: %v",
			defaultPersistenceManagerId, err)
	}
	mainPersistenceManager = persistenceManager
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func sendErrorResponse(writer http.ResponseWriter, message string, code int) {
	response, err := errorResponse(message)
	if err != nil {
		http.Error(writer, "An error occurred", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(code)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func errorResponse(message string) ([]byte, error) {
	errorResponse := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: false,
		Message: message,
	}

	return json.Marshal(errorResponse)
}

func sendSuccessJsonResponse(writer http.ResponseWriter, content interface{}) {
	response, err := json.Marshal(content)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
