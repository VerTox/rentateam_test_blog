package post_list

import (
	"errors"
	"github.com/VerTox/rentateam_test_blog/domain/model"
)

type Response struct {
	Posts []*model.Post
}

type Repositories struct {
	Post model.PostRepository
}

func Run(r *Repositories) (*Response, error) {
	if r == nil || r.Post == nil {
		return nil, errors.New("invalid case initialization")
	}

	p, err := r.Post.GetList()

	if err != nil {
		return nil, err
	}

	return &Response{Posts: p}, nil
}
