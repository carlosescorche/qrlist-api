package handlers

import (
	"net/http"

	"github.com/carlosescorche/qrlist-api/server"
	"github.com/carlosescorche/qrlist-api/utils/api"
	"github.com/gorilla/mux"
)

func HandlerListObserve(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		listId := vars["id"]

		socket, err := server.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Error upgrading connection", http.StatusInternalServerError)
			return
		}

		client := server.NewClient(listId, s.Hub, socket)
		s.Hub.Register <- client

		go client.Write()

		api.Success(w, nil, http.StatusOK)
	}
}
