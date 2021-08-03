package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ria-mehra16/Golang-Cassandra/pkg/adding"
	"github.com/Ria-mehra16/Golang-Cassandra/pkg/deleting"
	"github.com/Ria-mehra16/Golang-Cassandra/pkg/http/rest"
	"github.com/Ria-mehra16/Golang-Cassandra/pkg/reading"
	"github.com/Ria-mehra16/Golang-Cassandra/pkg/storage"
)

func main() {

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}

	rs := reading.NewService(r)
	as := adding.NewService(r)
	dl := deleting.NewService(r)

	fmt.Println("starting server on port 8080...")
	router := rest.InitHandlers(rs, as, dl)
	log.Fatal(http.ListenAndServe(":8080", router))
}
