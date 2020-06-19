package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oinume/todomvc/backend/repository"

	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.opencensus.io/plugin/ochttp"
	"go.uber.org/zap"

	"github.com/oinume/todomvc/backend/model"
	"github.com/oinume/todomvc/proto-gen/go/proto/todomvc"
)

type TodoItemsStore struct {
	items map[string]*todomvc.TodoItem
}

func (store *TodoItemsStore) Save(item *todomvc.TodoItem) error {
	store.items[item.Id] = item
	return nil
}

func (store *TodoItemsStore) Load(id string) (*todomvc.TodoItem, error) {
	if item, ok := store.items[id]; ok {
		return item, nil
	}
	return nil, fmt.Errorf("cannot find TodoItem for %s", id)
}

type server struct {
	todoRepo    repository.TodoRepository
	logger      *zap.Logger
	store       *TodoItemsStore
	unmarshaler *jsonpb.Unmarshaler
}

func NewServer(todoRepo repository.TodoRepository, logger *zap.Logger) *server {
	return &server{
		todoRepo: todoRepo,
		logger:   logger,
		store: &TodoItemsStore{
			items: make(map[string]*todomvc.TodoItem, 100),
		},
		unmarshaler: &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		},
	}
}

func (s *server) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(accessLogMiddleware(s.logger))
	r.Handle("/todos", ochttp.WithRouteTag(http.HandlerFunc(s.CreateTodo), "/todos")).Methods("POST")

	//r.HandleFunc("/todos", s.fetcher).Methods("GET")
	//r.HandleFunc("/todos/{id}", s.fetcher).Methods("Put")
	return r
}

func (s *server) CreateTodo(w http.ResponseWriter, r *http.Request) {
	req := &todomvc.CreateTodoRequest{}
	if err := s.unmarshaler.Unmarshal(r.Body, req); err != nil {
		internalServerError(s.logger, w, err)
		return
	}

	id := uuid.New().String()
	item := &todomvc.TodoItem{
		Id:    id,
		Title: req.Title,
	}
	if err := s.store.Save(item); err != nil {
		internalServerError(s.logger, w, err)
		return
	}

	todo := &model.Todo{
		ID:    id,
		Title: req.Title,
	}
	if err := s.todoRepo.Create(r.Context(), todo); err != nil {
		internalServerError(s.logger, w, err)
		return
	}

	writeJSON(w, http.StatusCreated, item)
}

func internalServerError(logger *zap.Logger, w http.ResponseWriter, err error) {
	//switch _ := errors.Cause(err).(type) { // TODO:
	//default:
	// unknown error
	//sUserID := ""
	//if userID == 0 {
	//	sUserID = fmt.Sprint(userID)
	//}
	//util.SendErrorToRollbar(err, sUserID)
	//fields := []zapcore.Field{
	//	zap.Error(err),
	//}
	//if e, ok := err.(errors.StackTracer); ok {
	//	b := &bytes.Buffer{}
	//	for _, f := range e.StackTrace() {
	//		fmt.Fprintf(b, "%+v\n", f)
	//	}
	//	fields = append(fields, zap.String("stacktrace", b.String()))
	//}
	//if appLogger != nil {
	//	appLogger.Error("internalServerError", fields...)
	//}

	logger.Error("caught error", zap.Error(err))
	http.Error(w, fmt.Sprintf("Internal server Error\n\n%v", err), http.StatusInternalServerError)
	//if !config.IsProductionEnv() {
	//	fmt.Fprintf(w, "----- stacktrace -----\n")
	//	if e, ok := err.(errors.StackTracer); ok {
	//		for _, f := range e.StackTrace() {
	//			fmt.Fprintf(w, "%+v\n", f)
	//		}
	//	}
	//}
}

func writeJSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, `{ "status": "Failed to Encode as writeJSON" }`, http.StatusInternalServerError)
	}
}

//func writeHTML(w http.ResponseWriter, code int, body string) {
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	w.WriteHeader(code)
//	if _, err := fmt.Fprint(w, body); err != nil {
//		http.Error(w, "Failed to write HTML", http.StatusInternalServerError)
//	}
//}

//func writeHTMLWithTemplate(
//	w http.ResponseWriter,
//	code int,
//	t *template.Template,
//	data interface{},
//) {
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	w.WriteHeader(code)
//	if err := t.Execute(w, data); err != nil {
//		internalServerError(w, err)
//	}
//}