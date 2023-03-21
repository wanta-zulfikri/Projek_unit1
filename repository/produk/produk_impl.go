package produk

import (
	"database/sql"
	"errors"


	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

type Produk struct {
	db *sql.DB
}

func InitProduk(db *sql.DB) ProdukInterface {
	return &Produk{db}
}

func (p *Produk) Tambahproduk(data *entities.Produk, userID int) error {
	row,err := p.db.Exec("INSERT INTO Produk (nama_produk, user_id) values(?, ?)", data.Nama_produk, userID)
	if err != nil {
		return errors.New("gagal membuat produk")
	}


	rowaff,err := row.RowsAffected()
	if err != nil {
		return errors.New("gagal membuat produk")
	}
	if rowaff > 0 {
		return nil

	} 
	return errors.New("gagal membuat produk") 
}

func (p *Produk) Getbynama(nama string)(*entities.Produk, error) {
	res := &entities.Produk{}
	row := p.db.QueryRow("select id, nama_produk, user_id from Produk where nama_produk = ? and deleted_at is null",nama)
	if row.Err() != nil {
		return nil,errors.New("nama produk tidak terdaftar")
	}
	err:= row.Scan(&res.Id, &res.Nama_produk, &res.User_id)
	if err != nil {
		return nil,errors.New("nama produk tidak terdaftar")
	}
	return res, nil
}
