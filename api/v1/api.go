package v1

import (
	"github.com/VerTox/rentateam_test_blog/api"
	"github.com/VerTox/rentateam_test_blog/domain"
	"github.com/gorilla/mux"
)

const RoutePrefix = "/api/v1"

type ApiV1 struct {
	*api.Api
}

func NewV1(r *mux.Router, c *domain.Context) {
	a := &ApiV1{
		Api: &api.Api{
			Router:  r.PathPrefix(RoutePrefix).Subrouter(),
			Context: c,
		},
	}

	a.bootRoutes()
}

func (a *ApiV1) bootRoutes() {
	r := a.Router

	r.HandleFunc("/posts", a.RouteWrapper(a.GetPostList)).Methods("GET")
	r.HandleFunc("/posts", a.RouteWrapper(a.CreatePost)).Methods("POST")
}
