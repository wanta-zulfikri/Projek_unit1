package migration

import (
	"github.com/wanta-zulfikri/Projek_unit1/db"
	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func Migration() {
	db := db.InitDb()
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS user (
		id int not null auto_increment,
		user_name VARCHAR (50) not null, 
		password varchar(50) not null,
		deleted_at timestamp default null,
		role varchar (50) not null,
		primary key(id)
		);`)
	helper.PanicIfError(err)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS produk(
		id INT not null auto_increment PRIMARY key ,
		nama_produk VARCHAR(50) not null,
		user_id INT (13) not null,
		deleted_at varchar(50) default null,
		updated_by varchar (50) default null,
		FOREIGN KEY (user_id) REFERENCES user(id)
		);`)
	_, err = db.Exec(`create table IF NOT EXISTS customer (
				id int not null auto_increment primary key,
				nama varchar(50) not null,
				alamat varchar(255) not null,
				deleted_at timestamp default null
				);`)
	helper.PanicIfError(err)
	_, err = db.Exec(`create table IF NOT EXISTS transaksi (
		id int primary key not null auto_increment,
		user_id int not null,
		tgl_transaksi timestamp not null,
		customer_id int  not null,
		deleted_at timestamp default null,  
		foreign key (user_id) references user(id),
		foreign key (customer_id) references customer(id)
		);`)
	helper.PanicIfError(err)

	_, err = db.Exec(`create table IF NOT EXISTS transaksi_item(
		id int not null auto_increment primary key ,
		qty int not null,
		price int not null,
		transaksi_id int not null, 
		produk_id int not null,
		foreign key (transaksi_id) references transaksi(id),
		foreign key (produk_id) references produk(id)
		
	);`)
	helper.PanicIfError(err)
	defer db.Close()

}
