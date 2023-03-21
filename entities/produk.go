package entities

import (
	"database/sql"


)
type Produk struct {
	Id int
	Nama_produk string
	User_id string
	Qty int
	Nama_pembuat string 
	Nama_Pengganti sql.NullString
}
