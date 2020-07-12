package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kopdar/kopdar-backend/internal/user/endpoint"
	httplib "github.com/kopdar/kopdar-backend/pkg/http"
)

// RegisterRoutesFunc represents the function for registering routes to a chi.Router.
type RegisterRoutesFunc func(r chi.Router)

func NewRegisterRoutesFunc(endpoints *endpoint.Set) RegisterRoutesFunc {
	return func(r chi.Router) {
		options := []httptransport.ServerOption{
			httptransport.ServerErrorEncoder(httplib.NewErrorHandler()),
		}
		r.Group(func(r chi.Router) {
			// TODO: ADD SOME MIDDLEWARE NEEDED
			r.Get("/user", httptransport.NewServer(
				endpoints.MakeFindAllEndpoint,
				func(ctx context.Context, r *http.Request) (interface{}, error) {
					return "", nil
				},
				httplib.ShowResponse(),
				options...).ServeHTTP)
		})
	}
}
