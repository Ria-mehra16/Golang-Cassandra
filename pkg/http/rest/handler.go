package rest

import (
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
