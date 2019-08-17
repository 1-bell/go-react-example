package db

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=./mocks/mock_db.go -package=db github.com/bogdanguranda/go-react-example/db DB

// DB describes CRUD methods for a database managing persons.
type DB interface {
	CreatePerson(person *Person) error
	DeletePerson(email string) error
	ListPersons(sortBy string) ([]*Person, error)

	Close()
}

// MySqlDB implements DB interface using MySQL as database.
type MySqlDB struct {
	db *sql.DB
}

// NewMySqlDB creates a new MySqlDB.
func NewMySqlDB(mySqlURL string) (*MySqlDB, error) {
	db, err := TryConnect(mySqlURL, 3, 5)
	if err != nil {
		return nil, err
	}

	return &MySqlDB{db: db}, nil
}

// TryConnect retry system for connecting to MySQL database.
func TryConnect(dsn string, delay, retries int) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	for ; err != nil && retries > 0; retries-- {
		time.Sleep(time.Second * time.Duration(delay))
		db, err = sql.Open("mysql", dsn)
	}
	return db, err
}

// Close closes the database connection.
func (my *MySqlDB) Close() {
	if err := my.db.Close(); err != nil {
		logrus.Panic("Failed to close MySQL db.")
	}
}

// CreatePerson creates a new person in the DB.
func (my *MySqlDB) CreatePerson(person *Person) error {
	insertQuery, err := my.db.Prepare("INSERT INTO Persons VALUES(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	if _, err := insertQuery.Exec(person.Name, person.Age, person.Balance, person.Email, person.Address); err != nil {
		return err
	}

	return nil
}

// DeletePerson deletes a person from the DB.
func (my *MySqlDB) DeletePerson(email string) error {
	updateQuery, err := my.db.Prepare("DELETE FROM Persons WHERE Email = ?;")
	if err != nil {
		return err
	}

	if _, err = updateQuery.Exec(email); err != nil {
		return err
	}

	return nil
}

// ListPersons retrieves all persons sorting by the given column.
func (my *MySqlDB) ListPersons(sortBy string) ([]*Person, error) {
	if sortBy != "name" || sortBy != "email" {
		return nil, errors.New("Unsupported sorting column: " + sortBy)
	}

	rows, err := my.db.Query("SELECT * FROM Persons ORDER BY " + sortBy + " DESC;")
	if err != nil {
		return nil, err
	}

	persons := []*Person{}
	for rows.Next() {
		person := &Person{}
		if err := rows.Scan(&person.Name, &person.Age, &person.Balance, &person.Email, &person.Address); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	return persons, nil
}
