package produk


import "github.com/wanta-zulfikri/Projek_unit1/entities"

type ProdukInterface interface {
	Tambahproduk(data *entities.Produk, userID int) error
	Getbynama(nama string)(*entities.Produk, error)
	GetAllProduk(offset int) ([]*entities.Produk, error)
	GetAll() ([]*entities.Produk, error)
	UpdateProduk(data *entities.Produk, user_name string) error
	Delete(userid int) error

}

