package api

import (
	"Comments/pkg/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	db     db.Storage
	router *mux.Router
}

func New(db db.Storage) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}
func (api *API) endpoints() {
	api.router.HandleFunc("/addComment", api.addComments).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/posts", api.comments).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) addComments(w http.ResponseWriter, r *http.Request) {
	var comm db.Comment
	err := json.NewDecoder(r.Body).Decode(&comm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.addComment(comm)
	w.WriteHeader(http.StatusOK)
}

func (api *API) comments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
