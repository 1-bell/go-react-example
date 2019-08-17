package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.org/bogdanguranda/go-react-example/db"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// mapPersonPayload maps a http request to create a new person.
func (dAPI *DefaultAPI) mapPersonPayload(r *http.Request) (*db.Person, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	person := db.Person{}
	return &creds, json.Unmarshal(body, &person)
}
