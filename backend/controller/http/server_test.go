package http

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/oinume/todomvc/backend/config"
	"github.com/oinume/todomvc/backend/infrastructure/mysql"
	"github.com/oinume/todomvc/backend/repository"
)

var (
	db       *sql.DB
	todoRepo repository.TodoRepository
)

func TestMain(m *testing.M) {
	// Overwrite database name for unit test
	if err := os.Setenv("MYSQL_DATABASE", "todomvc_test"); err != nil {
		panic(err)
	}
	config.MustProcessDefault()
	ldb, err := mysql.NewDB(config.DefaultVars.DBURL())
	if err != nil {
		panic("Failed to mysql.NewDB: " + err.Error())
	}

	db = ldb
	todoRepo = mysql.NewTodoRepository(db)
	os.Exit(m.Run())
}
