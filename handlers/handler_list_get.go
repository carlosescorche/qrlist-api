package handlers

import (
	"net/http"

	"github.com/carlosescorche/qrlist/model/list"
	"github.com/carlosescorche/qrlist/model/subscription"
	"github.com/carlosescorche/qrlist/server"
	"github.com/carlosescorche/qrlist/utils/api"
	e "github.com/carlosescorche/qrlist/utils/errors"
	"github.com/gorilla/mux"
)

type output struct {
	list *list.List
	subs []subscription.Subscription
}

func HandlerListGet(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		list, err := list.FindById(id)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusBadRequest)
			return
		}

		api.Success(w, list, http.StatusOK)
	}
}
