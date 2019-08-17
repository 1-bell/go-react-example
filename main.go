package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.org/bogdanguranda/go-react-example/api"
	"github.org/bogdanguranda/go-react-example/db"
)

const (
	portAPI = "8080"
)

func main() {
	dbMySQL := createDBConnection()
	defer dbMySQL.Close()

	startAPIServer(dbMySQL)
}

func createDBConnection() db.DB {
	mySqlURL := os.Getenv("MYSQL_URL")
	if mySqlURL == "" {
		logrus.Fatal("Env var MYSQL_URL  was not set!")
	}

	dbMySQL, err := db.NewMySqlDB(mySqlURL)
	if err != nil {
		logrus.Fatal(errors.Wrapf(err, "failed to start MySQL"))
	}

	return dbMySQL
}

func startAPIServer(dbMySQL db.DB) {
	appAPI := api.NewDefaultAPI(dbMySQL)

	router := mux.NewRouter()
	router.HandleFunc("/app/people", appAPI.Register).Methods(http.MethodGet)
	router.HandleFunc("/app/people", appAPI.LogIn).Methods(http.MethodPost)
	router.HandleFunc("/app/people", appAPI.LogOut).Methods(http.MethodDelete)

	logrus.Info("REST API server listening on port " + portAPI)
	if err := http.ListenAndServe(":"+portAPI, router); err != nil {
		logrus.Fatal("Failed to listen and serve on port " + portAPI)
	}
}
