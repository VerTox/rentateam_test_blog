package post_create

import (
	"errors"
	"github.com/VerTox/rentateam_test_blog/domain/model"
)

type Request struct {
	Post *model.Post
}

type Response struct {
	Post *model.Post
}

type Repositories struct {
	Post model.PostRepository
}

func Run(r *Repositories, request *Request) (*Response, error) {
	if r == nil || r.Post == nil || request == nil || request.Post == nil {
		return nil, errors.New("invalid post_create case initialization")
	}

	err := request.Post.Valid()

	if err != nil {
		return nil, err
	}

	p, err := r.Post.Store(request.Post)

	if err != nil {
		return nil, err
	}

	return &Response{Post: p}, nil
}
