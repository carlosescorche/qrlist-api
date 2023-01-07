package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosescorche/qrlist-api/model/subscription"
	"github.com/carlosescorche/qrlist-api/server"
	"github.com/carlosescorche/qrlist-api/utils/api"
	e "github.com/carlosescorche/qrlist-api/utils/errors"
	"github.com/carlosescorche/qrlist-api/utils/validator"
	"github.com/gorilla/mux"
)

type Payload struct {
	Status string `json:"status" validate:"oneof='ACCEPTED' 'CANCELLED' 'FINISHED'"`
}

func HandlerSubscriptionUpdate(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var payload Payload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			api.Error(w, e.ErrInvalidPayload, http.StatusBadRequest)
			return
		}

		if errs, ok := validator.ValidateStruct(payload); !ok {
			api.Error(w, e.NewPayloadError(errs), http.StatusBadRequest)
			return
		}

		subs, err := subscription.FindById(id)
		if err != nil {
			err := e.NewCustomError("errBadRequest", "Subscription not found", nil)
			api.Error(w, err, http.StatusBadRequest)
			return
		}

		subs.Status = payload.Status

		err = subscription.Update(*subs)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusInternalServerError)
			return
		}

		s.Hub.Broadcast(map[string]interface{}{
			"type": "subscription_updated",
			"data": subs,
		}, subs.ListId.Hex())

		api.Success(w, nil, http.StatusAccepted)
	}
}
