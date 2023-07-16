package users

import (
	"context"
	"database/sql"
)

func CreateNewUser(ctx context.Context, db *sql.DB, payload CreateUserPayload) (err error) {

	userEntity := UserEntity{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Address:   payload.Address,
		Username:  payload.Username,
		Password:  payload.Password,
	}

	if _, err = SaveUserRepo(ctx, db, userEntity); err != nil {
		return
	}

	return
}
