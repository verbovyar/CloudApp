package handlers

import (
	"errors"
	"io"
	"net/http"
	"project/CloudApp/internal/repositories/interfaces"

	"github.com/gorilla/mux"
)

var errorNoSuchKey = errors.New("no such key")

type Handlers struct {
	store interfaces.RepoInterface
}

func New(store interfaces.RepoInterface) *Handlers {
	return &Handlers{store: store}
}

func (h *Handlers) PutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = h.store.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := h.store.Get(key)
	if errors.Is(err, errorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte(value))
}

func (h *Handlers) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := h.store.Delete(key)
	if errors.Is(err, errorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
