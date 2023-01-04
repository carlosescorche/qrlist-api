package handlers

import (
	"fmt"
	"net/http"

	"github.com/carlosescorche/qrlist/model/counter"
	"github.com/carlosescorche/qrlist/model/list"
	"github.com/carlosescorche/qrlist/model/subscription"
	"github.com/carlosescorche/qrlist/server"
	"github.com/carlosescorche/qrlist/utils/api"
	e "github.com/carlosescorche/qrlist/utils/errors"
	"github.com/gorilla/mux"
)

func HandlerListSignup(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		listId := vars["id"]

		list, err := list.FindById(listId)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusBadRequest)
			return
		}

		counter, err := counter.GetCounter(list.Id)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusInternalServerError)
			return
		}

		subs := subscription.NewSubscription()
		subs.ListId = list.Id
		subs.Number = fmt.Sprintf("%v-%v", "A", counter.Count)

		err = subscription.Insert(subs)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusInternalServerError)
		}

		s.Hub.Broadcast(map[string]interface{}{
			"type": "subscription_added",
			"data": subs,
		}, subs.ListId.Hex())

		api.Success(w, subs, http.StatusOK)
	}
}
