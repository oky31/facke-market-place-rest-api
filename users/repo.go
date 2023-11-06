package users

import (
	"context"
	"database/sql"
	"log"
)

type UserEntity struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Address   string
	Username  string
	Password  string
}

func SaveUserRepo(ctx context.Context, db *sql.DB, user UserEntity) (lastInsertId int64, err error) {

	queryInsert := "INSERT INTO users(first_name, last_name, email, address, username, password) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.ExecContext(ctx,
		queryInsert,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Address,
		user.Username,
		user.Password,
	)

	if err != nil {
		log.Print(err.Error())
		return
	}

	lastInsertId, err = result.LastInsertId()
	if err != nil {
		log.Print(err.Error())
		return
	}

	return
}

func FindUserRepo(ctx context.Context, db *sql.DB, id int64) (UserEntity, error) {

	user := UserEntity{}

	queryFind := "SELECT * FROM users WHERE users.id = ?"
	if err := db.QueryRowContext(ctx, queryFind, id).
		Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Username); err != nil {

		return user, err
	}

	return user, nil
}
