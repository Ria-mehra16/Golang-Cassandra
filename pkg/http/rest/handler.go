package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Ria-mehra16/Golang-Cassandra/pkg/adding"
	"github.com/Ria-mehra16/Golang-Cassandra/pkg/deleting"
	"github.com/Ria-mehra16/Golang-Cassandra/pkg/reading"
	"github.com/gorilla/mux"
)

func InitHandlers(rs reading.Service, as adding.Service, dl deleting.Service) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies", getAllCandiesHandler(rs)).Methods("GET")

	//Adding
	router.HandleFunc("/api/candy", addCandy(as)).Methods("POST")
	router.HandleFunc("/api/deletecandy", deleteCandy(dl)).Methods("DELETE")
	return router
}

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to our candy shop!")
	}
}

func getAllCandiesHandler(rs reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := rs.GetAllCandyNames()
		if err != nil {
			http.Error(w, "Cannot process your request at this time...", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(cs)
	}
}

func deleteCandy(dl deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var nc deleting.Candy
		if err := json.NewDecoder(r.Body).Decode(&nc); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		id, err := dl.DeleteCandy(nc)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		nc.Id = id
		json.NewEncoder(w).Encode(nc)
	}
}

func addCandy(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var nc adding.Candy
		if err := json.NewDecoder(r.Body).Decode(&nc); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		id, err := as.AddCandy(nc)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		nc.Id = id
		json.NewEncoder(w).Encode(nc)
	}
}
