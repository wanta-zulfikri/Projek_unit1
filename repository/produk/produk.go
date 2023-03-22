package produk

import "github.com/wanta-zulfikri/Projek_unit1/entities"

type ProdukInterface interface {
	Tambahproduk(data *entities.Produk, userID int) error
	Getbynama(nama string) (*entities.Produk, error)
	GetAllProduk(offset int) ([]*entities.Produk, error)
	GetAll() ([]*entities.Produk, error)
	UpdateProduk(data *entities.Produk, user_name string) error
	UpdateStok(newstok int, produkid int) error
	Delete(userid int) error 
	Harga(data *entities.Produk, harga int) error
	HargaProduk(harga int) (*entities.Produk, error)
	
}
