package transaksi

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

type Transaksi struct {
	db *sql.DB
}

func InitTransaksi(db *sql.DB) TransaksiInterface {
	return &Transaksi{db}
}

func (trx *Transaksi) Create(data *entities.Transaksi) (error, int) {

	row, err := trx.db.Exec("INSERT INTO transaksi(user_id,tgl_transaksi,customer_id) VALUES(?,?,?)", data.UserId, time.Now(), data.CustomerId)
	if err != nil {
		return errors.New("Transaksi Tidak Berhasil Ditambahkan"), 0
	}
	rowaff, _ := row.RowsAffected()
	id, _ := row.LastInsertId()
	if rowaff > 0 {
		return nil, int(id)
	}
	return errors.New("Transaksi Tidak Berhasil Ditambahkan"), 0
}

func (trx *Transaksi) GetWithLimit(offset int) ([]*entities.Transaksi, error) {
	res := []*entities.Transaksi{}
	rows, err := trx.db.Query(fmt.Sprintf("SELECT id,user_id,tgl_transaksi,customer_id, FROM transaksi WHERE deleted_at IS NULL LIMIT %d OFFSET %d", config.LimitPage, offset))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := &entities.Transaksi{}
		err := rows.Scan(&row.Id, &row.UserId, &row.Tanggal, &row.CustomerId)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func (trx *Transaksi) GetAll() ([]*entities.Transaksi, error) {
	res := []*entities.Transaksi{}
	rows, err := trx.db.Query("SELECT id,user_id,tgl_transaksi,customer_id, FROM transaksi WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := &entities.Transaksi{}
		err := rows.Scan(&row.Id, &row.UserId, &row.Tanggal, &row.CustomerId)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func (trx *Transaksi) InsertItems(data *entities.TransaksiItem) error {
	row, err := trx.db.Exec("INSERT INTO transaksi_item(qty,price,transaksi_id,produk_id)VALUES(?,?,?,?)", data.Qty, data.Price, data.TrxId, data.ProdukId)
	if err != nil {
		return errors.New("Transaksi Tidak Berhasil Ditambahkan")
	}

	rowaff, _ := row.RowsAffected()
	if rowaff > 0 {
		return nil
	}
	return errors.New("Transaksi Tidak Berhasil Ditambahkan")
}
