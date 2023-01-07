package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosescorche/qrlist-api/model/list"
	"github.com/carlosescorche/qrlist-api/server"
	"github.com/carlosescorche/qrlist-api/utils/api"
	e "github.com/carlosescorche/qrlist-api/utils/errors"
	"github.com/carlosescorche/qrlist-api/utils/validator"
)

type HandlerListCreatePayload struct {
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=350"`
	Code        string `json:"code" validate:"required,max=5"`
}

func HandlerListCreate(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload HandlerListCreatePayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			api.Error(w, e.ErrInvalidPayload, http.StatusBadRequest)
			return
		}

		if errs, ok := validator.ValidateStruct(payload); !ok {
			api.Error(w, e.NewPayloadError(errs), http.StatusBadRequest)
			return
		}

		newList := list.NewList()
		newList.Name = payload.Name
		newList.Description = payload.Description

		err = list.Insert(newList)
		if err != nil {
			api.Error(w, e.ErrInternal, http.StatusInternalServerError)
			return
		}

		api.Success(w, newList, http.StatusCreated)
	}
}
