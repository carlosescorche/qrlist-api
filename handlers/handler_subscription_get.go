package handlers

import (
	"net/http"

	"github.com/carlosescorche/qrlist-api/model/subscription"
	"github.com/carlosescorche/qrlist-api/server"
	"github.com/carlosescorche/qrlist-api/utils/api"
	e "github.com/carlosescorche/qrlist-api/utils/errors"
	"github.com/gorilla/mux"
)

func HandlerSubscriptionGet(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		subs, err := subscription.FindById(id)
		if err != nil {
			err := e.NewCustomError("errBadRequest", "Subscription not found", nil)
			api.Error(w, err, http.StatusBadRequest)
			return
		}

		api.Success(w, subs, http.StatusOK)
	}
}
