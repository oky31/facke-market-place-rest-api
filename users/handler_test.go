package users

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		panic(err)
	}

	return db
}

func createTableUsers(db *sql.DB) {
	queryCreateTableUsers := `
		CREATE TABLE users (
			id          INT AUTO_INCREMENT,
			first_name  VARCHAR(50)  NOT NULL,
			last_name   VARCHAR(50)  NULL,
			email       VARCHAR(100) NOT NULL,
			address     VARCHAR(255) NOT NULL,
			username    VARCHAR(50)  NOT NULL,
			password    VARCHAR(255) NOT NULL,
			created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY(id)
		)`
	result, err := db.Exec(queryCreateTableUsers)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {

	db := DbConnection()
	createTableUsers(db)

	payload := CreateUserPayload{
		FirstName: "Oky",
		LastName:  "Saputra",
		Email:     "oky@gmail.com",
		Address:   "Jl. Perjalanan panjang",
		Username:  "oky55",
		Password:  "12345",
	}

	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("Error in marshal json %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/user/create", bytes.NewReader(bytesPayload))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	createUserHandler := CreateUserHandler{db: db}
	createUserHandler.ServeHTTP(res, req)

	assertHttpStatus(t, http.StatusCreated, res.Code)
}
