package users

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

func DbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
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

func dropTable(db *sql.DB) {
	db.Exec("DROP TABLE users")
}

func payloadSuccessCreateUser() CreateUserPayload {
	return CreateUserPayload{
		FirstName: "Oky",
		LastName:  "Saputra",
		Email:     "oky@gmail.com",
		Address:   "Jl. Perjalanan panjang",
		Username:  "oky55",
		Password:  "12345",
	}

}

func TestCreateUser(t *testing.T) {
	db := DbConnection()
	createTableUsers(db)

	payload := payloadSuccessCreateUser()

	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("Error in marshal json %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/users/create", bytes.NewReader(bytesPayload))
	req.Header.Set("Content-Type", ContentTypeJson)

	res := httptest.NewRecorder()

	logJsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJsonHandler)
	createUserHandler := CreateUserHandler{db: db, logger: logger}
	createUserHandler.ServeHTTP(res, req)

	assertHttpStatus(t, http.StatusCreated, res.Code)
	assertContentTypeJson(t, res.Header().Get("Content-Type"))

	dropTable(db)
}

func assertContentTypeJson(t testing.TB, contentTypeActual string) {
	t.Helper()
	if contentTypeActual != ContentTypeJson {
		t.Errorf("expected content type %v actual %v", ContentTypeJson, contentTypeActual)
	}
}
