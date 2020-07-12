package http

import (
	"context"
	"encoding/json"
	"github.com/kopdar/kopdar-backend/pkg/errors"
	"net/http"
)

func NewErrorHandler() func(ctx context.Context, err error, w http.ResponseWriter) {
	return func(_ context.Context, err error, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(getHttpCodeFromErr(err))
		res := response{}
		res.Code = getCodeFromErr(err)
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
	}
}

func getHttpCodeFromErr(err error) int {
	if errors.Is(errors.RequestFailed, err) {
		return http.StatusBadRequest
	}

	if errors.Is(errors.NotFound, err) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}

func getCodeFromErr(err error) string {
	if v, ok := err.(*errors.Error); ok {
		return v.Code
	}
	return errors.CodeInternalError
}