// modeltest provides functions to create database fixtures for unit test.
package modeltest

import (
	"database/sql"
	mrand "math/rand"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/oinume/todomvc/backend/model"
)

var (
	letters = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
	random  = mrand.New(mrand.NewSource(time.Now().UnixNano()))
)

// NewTodo returns `*model.Todo` with setters function to specify field value for *model.Todo.
func NewTodo(setters ...func(*model.Todo)) *model.Todo {
	todo := &model.Todo{}
	for _, setter := range setters {
		setter(todo)
	}
	if todo.ID == "" {
		todo.ID = uuid.New().String()
	}
	if todo.Title == "" {
		todo.Title = RandomString(10)
	}
	return todo
}

// RandomString creates random string of alphabet and number with given length.
func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[random.Intn(len(letters))]
	}
	return string(b)
}

// TruncateAllTables truncates all tables of the database.
// You need to call this function to remove all rows to ensure your test starts with clean database.
func TruncateAllTables(t *testing.T, db *sql.DB, dbName string) {
	t.Helper()

	tables, err := LoadAllTables(t, db, dbName)
	if err != nil {
		t.Fatal(err)
	}
	for _, table := range tables {
		_, err := db.Exec("TRUNCATE TABLE " + table)
		if err != nil {
			t.Fatalf("failed to truncate table: table=%v, err=%v", table, err)
		}
	}
}

func LoadAllTables(t *testing.T, db *sql.DB, dbName string) ([]string, error) {
	t.Helper()

	sql := "SELECT table_name AS name FROM information_schema.tables WHERE table_schema = ?"
	rows, err := db.Query(sql, dbName)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	names := make([]string, 0, 10)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		if name == "goose_db_version" {
			continue
		}
		names = append(names, name)
	}

	return names, nil
}
