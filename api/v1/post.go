package v1

import (
	"encoding/json"
	"github.com/VerTox/rentateam_test_blog/api"
	"github.com/VerTox/rentateam_test_blog/domain/cases/post_create"
	"github.com/VerTox/rentateam_test_blog/domain/cases/post_list"
	"github.com/VerTox/rentateam_test_blog/domain/model"
	"net/http"
)

func (a *ApiV1) GetPostList(c *api.Context, w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	resp, err := post_list.Run(&post_list.Repositories{Post: a.Context.Connection.Post()})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, ListResponse{
		Items: resp.Posts,
		Total: len(resp.Posts),
	}, nil
}

func (a *ApiV1) CreatePost(c *api.Context, w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	p := &model.Post{}

	err := json.Unmarshal(c.Body, p)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	resp, err := post_create.Run(&post_create.Repositories{Post: a.Context.Connection.Post()}, &post_create.Request{Post: p})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, resp.Post, nil
}
