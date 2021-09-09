package pg

import (
	"database/sql"
	"go-jwt/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) ExistUser(name string) bool {
	sqlQuery := `
	SELECT exists (
		SELECT id
		FROM PUBLIC.user
		WHERE name = $1
	)
`
	var exist bool
	err := u.DB.QueryRow(sqlQuery, name).Scan(&exist)

	if err != nil {
		return true
	}
	return exist
}

func (u *UserModel) Create(user models.SignUp) (string, error) {
	sqlQuery := `INSERT INTO public.user (name, password) values ($1, $2) returning id`
	var id string

	err := u.DB.QueryRow(sqlQuery, user.Name, user.Password).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
