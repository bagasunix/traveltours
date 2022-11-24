package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
)

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {

	responseBuilder := responses.NewErrorBuilder()
	if strings.Contains(err.Error(), errors.ERR_NOT_AUTHORIZED) {
		w.WriteHeader(http.StatusUnauthorized)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}
	if strings.Contains(err.Error(), errors.ERR_INVALID_KEY) {
		w.WriteHeader(http.StatusBadRequest)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(err.Error(), errors.ERR_NOT_FOUND) {
		w.WriteHeader(http.StatusNotFound)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(err.Error(), errors.ERR_ALREADY_EXISTS) {
		w.WriteHeader(http.StatusConflict)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(err.Error(), errors.ERR_UNMARSHAL) {
		w.WriteHeader(http.StatusBadRequest)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	errors.HandlerReturnedVoid(err)
	w.WriteHeader(http.StatusInternalServerError)
	responseBuilder.SetError(errors.CustomError(errors.ERR_SOMETHING_WRONG))
	json.NewEncoder(w).Encode(responseBuilder.Build())
}
