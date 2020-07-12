package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//ShowResponse generate a httptransport encode function to encode response to json
func ShowResponse() func(ctx context.Context, w http.ResponseWriter, data interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, data interface{}) error {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		res := response{}
		res.Code = "success"
		res.Message = "Proses Berhasil"
		res.Data = data
		return json.NewEncoder(w).Encode(res)
	}
}
