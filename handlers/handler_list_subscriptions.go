package handlers

import (
	"net/http"

	"github.com/carlosescorche/qrlist-api/model/list"
	"github.com/carlosescorche/qrlist-api/model/subscription"
	"github.com/carlosescorche/qrlist-api/server"
	"github.com/carlosescorche/qrlist-api/utils/api"
	e "github.com/carlosescorche/qrlist-api/utils/errors"
	"github.com/gorilla/mux"
)

func HandlerListSubscriptions(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		_, err := list.FindById(id)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusBadRequest)
			return
		}

		subs, err := subscription.FindByListId(id)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusBadRequest)
			return
		}

		api.Success(w, subs, http.StatusOK)
	}
}
