package produk

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

type Produk struct {
	db *sql.DB
}

func InitProduk(db *sql.DB) ProdukInterface {
	return &Produk{db}
}

func (p *Produk) Tambahproduk(data *entities.Produk, userID int) error {
	row, err := p.db.Exec("INSERT INTO Produk (nama_produk, user_id, qty) values(?, ?, ?)", data.Nama_produk, userID, data.Qty)
	fmt.Println(err)
	if err != nil {
		return errors.New("gagal membuat produk")
	}

	rowaff, err := row.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return errors.New("gagal membuat produk")
	}
	if rowaff > 0 {
		return nil

	}
	return errors.New("gagal membuat produk")
}

func (p *Produk) Getbynama(nama string) (*entities.Produk, error) {
	res := &entities.Produk{}
	row := p.db.QueryRow("select id,nama_produk,user_id,qty from Produk where nama_produk=? and deleted_at is null", nama)
	if row.Err() != nil {
		return nil, errors.New("nama produk tidak terdaftar")
	}
	err := row.Scan(&res.Id, &res.Nama_produk, &res.User_id)
	if err != nil {
		return nil, errors.New("nama produk tidak terdaftar")
	}
	return res, nil
}

func (p *Produk) UpdateProduk(data *entities.Produk, user_name string) error {
	res, err := p.db.Exec("UPDATE Produk SET qty=?,nama_produk=?,updated_by=? WHERE id=?", data.Qty, data.Nama_produk, user_name, data.Id)
	if err != nil {
		return err
	}
	resaff, _ := res.RowsAffected()
	if resaff > 0 {
		return nil
	}

	return errors.New("gagal memperbarui produk")
}

func (p *Produk) GetAll() ([]*entities.Produk, error) {
	res := []*entities.Produk{}
	rows, err := p.db.Query("select id, nama_produk, user_id, qty from Produk where deleted_at is null")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := &entities.Produk{}
		err := rows.Scan(&row.Id, &row.Nama_produk, &row.User_id, &row.Qty)
		if err != nil {
			return nil, err
		}
		res = append(res, row)

	}
	return res, nil
}

func (P *Produk) GetAllProduk(offset int) ([]*entities.Produk, error) {
	res := []*entities.Produk{}
	rows, err := P.db.Query(fmt.Sprintf("SELECT p.id,p.nama_produk,p.user_id,p.qty,u.user_name,p.updated_by from Produk p JOIN user u ON u.id=p.user_id where p.qty > 0 AND p.deleted_at is null limit %d offset %d", config.LimitPage, offset))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := &entities.Produk{}
		err := rows.Scan(&row.Id, &row.Nama_produk, &row.User_id, &row.Qty, &row.Nama_pembuat, &row.Nama_Pengganti)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func (p *Produk) GetAllow() ([]*entities.Produk, error) {
	res := []*entities.Produk{}
	rows, err := p.db.Query("SELECT id,nama_produk,user_id,qty from Produk where deleted_at is null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := &entities.Produk{}
		err := rows.Scan(&row.Id, &row.Nama_produk, &row.User_id, &row.Qty)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}
func (p *Produk) Delete(produkid int) error {
	res, err := p.db.Exec("UPDATE Produk set deleted_at=? WHERE Id=?", time.Now(), produkid)
	if err != nil {
		return err
	}
	resaff, _ := res.RowsAffected()
	if resaff > 0 {
		return nil
	}
	return errors.New("produk tidak berhasil dihapus")
}

func (p *Produk) UpdateStok(newstok int, produkid int) error {
	res, err := p.db.Exec("UPDATE Produk SET qty=? WHERE id=?", newstok, produkid)
	if err != nil {
		return err
	}
	resaff, _ := res.RowsAffected()
	if resaff > 0 {
		return nil
	}

	return errors.New("gagal memperbarui produk")
}
