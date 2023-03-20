package user

import (
	"database/sql"
	"errors"

	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

type User struct {
	db *sql.DB
}

func InitUser(db *sql.DB) UserInterface {
	return &User{db}
}

func (u *User) FindByUsername(username string) (*entities.User, error) {
	res := entities.User{}
	row := u.db.QueryRow("SELECT id,user_name,password,role FROM user WHERE user_name=? AND deleted_at IS NULL", username)
	if row.Err() != nil {
		return nil, errors.New("Username Tidak Terdaftar atau Akun Sudah Tidak Aktif")
	}
	row.Scan(&res.Id, &res.Username, &res.Password, &res.Role)
	return &res, nil
}
