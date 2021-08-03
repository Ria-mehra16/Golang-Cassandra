package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Ria-mehra16/Golang-Cassandra/pkg/deleting"
)

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
