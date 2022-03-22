package repository

import (
	"github.com/VerTox/rentateam_test_blog/domain"
	"github.com/VerTox/rentateam_test_blog/domain/model"
	"github.com/VerTox/rentateam_test_blog/repository/mysql"
)

type Connection struct {
	mysql *mysql.Connection
}

func New(m *mysql.Connection) domain.Connection {
	return &Connection{
		mysql: m,
	}
}

func (c *Connection) Post() model.PostRepository {
	return c.mysql.Post()
}

func (c *Connection) IsErrNotFound(err error) bool {
	return c.mysql.IsErrNotFound(err)
}
