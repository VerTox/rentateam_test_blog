package main

import (
	"fmt"
	"github.com/VerTox/rentateam_test_blog/repository"
	"github.com/VerTox/rentateam_test_blog/repository/mysql"
	"os"
)

func main() {
	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		panic("No APP_PORT provided")
	}

	dsn := os.Getenv("MYSQL_DSN")

	if dsn == "" {
		panic("No MYSQL_DSN provided")
	}

	m, err := mysql.NewConnection(dsn)

	if err != nil {
		panic(err)
	}

	err = m.Connection.AutoMigrate(
		mysql.Post{},
	).Error
	if err != nil {
		panic(err)
	}

	c := repository.New(m)

	a := &Application{
		Connection: c,
	}

	a.Run(fmt.Sprintf(":%s", appPort))
}
