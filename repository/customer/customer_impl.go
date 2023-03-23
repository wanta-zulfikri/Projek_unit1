package customer

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

type Customer struct {
	db *sql.DB
}

func InitCustomer(db *sql.DB) CustomerInterface {
	return &Customer{db}
}

func (cus *Customer) GetAll() ([]*entities.Customer, error) {
	res := []*entities.Customer{}
	rows, err := cus.db.Query("SELECT id,nama,alamat,no_hp from customer WHERE deleted_at IS NULL")
	if err != nil {
		return nil, errors.New("Gagal mengambil data")
	}
	defer rows.Close()
	for rows.Next() {
		row := &entities.Customer{}
		err := rows.Scan(&row.Id, &row.Nama, &row.Alamat, &row.NoHp)

		if err != nil {
			return nil, errors.New("Gagal Mengambil data")
		}
		res = append(res, row)
	}
	return res, nil

}
func (cus *Customer) Create(data *entities.Customer) error {

	row, err := cus.db.Exec("INSERT INTO customer(nama,alamat,no_hp) VALUES(?,?,?)", data.Nama, data.Alamat, data.NoHp)

	if err != nil {
		return errors.New("Data tidak berhasil dibuat")
	}
	rowaff, _ := row.RowsAffected()
	if rowaff > 0 {
		return nil
	}
	return errors.New("Data tidak berhasil dibuat")
}

func (cus *Customer) Update(data *entities.Customer) error {

	row, err := cus.db.Exec("UPDATE customer SET nama=?,alamat=?,no_hp=? WHERE id=?", data.Nama, data.Alamat, data.NoHp, data.Id)
	if err != nil {
		return errors.New("Data tidak berhasil dibuat")
	}
	rowaff, _ := row.RowsAffected()
	if rowaff > 0 {
		return nil
	}
	return errors.New("Data tidak berhasil dibuat")
}

func (cus *Customer) Delete(userid int) error {
	row, err := cus.db.Exec("UPDATE customer SET deleted_at=? WHERE id=?", time.Now(), userid)
	if err != nil {
		return errors.New("Data tidak berhasil dihapus")
	}
	rowaff, _ := row.RowsAffected()
	if rowaff > 0 {
		return nil
	}
	return errors.New("Data tidak berhasil hapus")
}

func (cus *Customer) FindByPhone(phonenum string) (*entities.Customer, error) {
	res := &entities.Customer{}
	row := cus.db.QueryRow("SELECT id,nama,alamat,no_hp from customer where no_hp=?", phonenum)
	if row.Err() != nil {
		return nil, errors.New("User belum terdaftar")
	}
	err := row.Scan(&res.Id, &res.Nama, &res.Alamat, &res.NoHp)
	if err != nil {
		return nil, errors.New("User belum terdaftar")
	}
	return res, nil
}

func (u *Customer) GetWithLimit(offset int) ([]*entities.Customer, error) {
	res := []*entities.Customer{}
	rows, err := u.db.Query(fmt.Sprintf("SELECT id,nama,alamat,no_hp from customer WHERE deleted_at IS NULL LIMIT %d OFFSET %d", config.LimitPage, offset))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := &entities.Customer{}
		err := rows.Scan(&row.Id, &row.Nama, &row.Alamat, &row.NoHp)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}
