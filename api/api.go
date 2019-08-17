package api

import (
	"encoding/json"
	"net/http"

	"github.org/bogdanguranda/go-react-example/db"
)

// API defines the REST API for managing persons.
type API interface {
	CreatePerson(w http.ResponseWriter, r *http.Request)
	DeletePerson(w http.ResponseWriter, r *http.Request)
	ListPersons(w http.ResponseWriter, r *http.Request)
}

// DefaultAPI default implementation of API.
type DefaultAPI struct {
	db db.DB
}

// NewDefaultAPI creates a new DefaultAPI.
func NewDefaultAPI(db db.DB) *DefaultAPI {
	return &DefaultAPI{db: db}
}

// CreatePerson handles creation of a new person.
func (dAPI *DefaultAPI) CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	person, err := dAPI.mapPersonPayload(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(Response{Error: "Invalid payload."})
		w.Write(resp)
		return
	}

	if err := dAPI.CreatePerson(person); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(Response{Error: "There was a problem with the server."})
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(SuccessResp{Cookie: cookie})
	w.Write(resp)
}

// DeletePerson handles removal of a person.
func (dAPI *DefaultAPI) DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	person, err := dAPI.mapPersonPayload(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(Response{Error: "Invalid payload."})
		w.Write(resp)
		return
	}

	if err := dAPI.db.DeletePerson(person.Email); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(Response{Error: "There was a problem with the server."})
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(profile)
	w.Write(resp)
}

// ListPersons handles retrieving and sorting persons.
func (dAPI *DefaultAPI) ListPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: get the query param sortBy here
	sortBy := "name"

	persons, err := dAPI.db.ListPersons(sortBy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(Response{Error: "There was a problem on the server."})
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(persons)
	w.Write(resp)
}
