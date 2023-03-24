package user

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

type User struct {
	db *sql.DB
}

func InitUser(db *sql.DB) UserInterface {
	return &User{db}
}

func (u *User) FindByUsername(username string) (*entities.User, error) {
	res := &entities.User{}
	row := u.db.QueryRow("SELECT id,user_name,password,role FROM user WHERE user_name=?", username)
	if row.Err() != nil {
		return nil, errors.New("Username Tidak Terdaftar")
	}
	err := row.Scan(&res.Id, &res.Username, &res.Password, &res.Role)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *User) GetAllByRole(role string) ([]*entities.User, error) {
	res := []*entities.User{}
	rows, err := u.db.Query("SELECT id,user_name,password,role FROM user WHERE role=? AND deleted_at IS NULL", role)
	if err != nil {
		return nil, err
	}
	rows.Next()
	for rows.Next() {
		row := &entities.User{}
		err := rows.Scan(&row.Id, &row.Username, &row.Password, &row.Role)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func (u *User) GetAllByRoleLimit(role string, offset int) ([]*entities.User, error) {
	res := []*entities.User{}
	rows, err := u.db.Query(fmt.Sprintf("SELECT id,user_name,password,role FROM user WHERE role=? AND deleted_at IS NULL LIMIT %d OFFSET %d", config.LimitPage, offset), role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := &entities.User{}
		err := rows.Scan(&row.Id, &row.Username, &row.Password, &row.Role)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func (u *User) GetAll() ([]*entities.User, error) {
	res := []*entities.User{}
	rows, err := u.db.Query("SELECT id,username,password,role from user where deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := &entities.User{}
		err := rows.Scan(&row.Id, &row.Username, &row.Password, &row.Role)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func (u *User) Create(data *entities.User) error {
	res, err := u.db.Exec("INSERT INTO user(user_name,password,role) VALUES(?,?,?)", data.Username, data.Password, data.Role)
	if err != nil {
		return err
	}
	rowaff, _ := res.RowsAffected()
	fmt.Println(rowaff, "ini err")
	if rowaff > 0 {
		return nil
	}
	return errors.New("User tidak berhasil dibuat")
}

func (u *User) Update(data *entities.User, userid int) error {
	res, err := u.db.Exec("UPDATE user set user_name=?,password=? WHERE id=?", data.Username, data.Password, userid)
	if err != nil {
		return err
	}
	resaff, _ := res.RowsAffected()
	if resaff > 0 {
		return nil
	}
	return errors.New("User tidak berhasil Update")
}
func (u *User) Delete(userid int) error {
	res, err := u.db.Exec("UPDATE user set deleted_at=? WHERE id=?", time.Now(), userid)
	if err != nil {
		return err
	}
	resaff, _ := res.RowsAffected()
	if resaff > 0 {
		return nil
	}
	return errors.New("User tidak berhasil dihapus")
}

func (u *User) InsertLog(oldusername, newusername string) error {
	res, err := u.db.Exec("INSERT INTO log_account(username_lama,username_baru,tanggal_perubahan) VALUES(?,?,?)", oldusername, newusername, time.Now())
	if err != nil {
		return err
	}
	resaff, _ := res.RowsAffected()
	if resaff > 0 {
		return nil
	}
	return errors.New("Log tidak berhasil ditambahkan")

}

func (u *User) GetLog() ([]*entities.Log, error) {
	res := []*entities.Log{}
	rows, err := u.db.Query("SELECT username_lama,username_baru,tanggal_perubahan FROM log_account")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := &entities.Log{}
		err := rows.Scan(&row.OldUsername, &row.NewUsername, &row.Date)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil

}

func (u *User) GetLogWithLimit(offset int) ([]*entities.Log, error) {

	res := []*entities.Log{}
	rows, err := u.db.Query(fmt.Sprintf("SELECT username_lama,username_baru,tanggal_perubahan FROM log_account LIMIT %d OFFSET %d", config.LimitPage, offset))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := &entities.Log{}
		err := rows.Scan(&row.OldUsername, &row.NewUsername, &row.Date)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}
