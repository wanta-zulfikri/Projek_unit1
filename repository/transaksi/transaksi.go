package transaksi

import "github.com/wanta-zulfikri/Projek_unit1/entities"

type TransaksiInterface interface {
	Create(data *entities.Transaksi) (error, int)
	InsertItems(data *entities.TransaksiItem) error
	GetWithLimit(offset int) ([]*entities.Transaksi, error)
	GetAll() ([]*entities.Transaksi, error)
	GetListItemByid(trxid int) ([]*entities.TransaksiItem, error)
	GetAllByUid(userid int) ([]*entities.Transaksi, error)
	GetWithLimitByUid(userid, offset int) ([]*entities.Transaksi, error)
	Delete(trxid int) error
}
