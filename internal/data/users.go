package data

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID       int64
	Isactive bool
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) Insert(user *User) error {

	query := `INSERT INTO users(id, isactive) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	_, err := u.DB.ExecContext(ctx, query, user.ID, user.Isactive)

	return err

}

func (u UserModel) Update(user *User) error {
	query := `UPDATE users SET isactive = $1 WHERE id = $2`

	ctx , cancel :=  context.WithTimeout(context.Background(),3* time.Second)

	defer cancel()

	_, err :=  u.DB.ExecContext(ctx, query, user.Isactive, user.ID)

	return err
}