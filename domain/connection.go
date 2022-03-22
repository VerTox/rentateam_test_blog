package domain

import "github.com/VerTox/rentateam_test_blog/domain/model"

type Connection interface {
	Post() model.PostRepository
	IsErrNotFound(err error) bool
}
