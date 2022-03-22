package api

import (
	"github.com/VerTox/rentateam_test_blog/domain"
	"github.com/gorilla/mux"
)

type Api struct {
	Router  *mux.Router
	Context *domain.Context
}
