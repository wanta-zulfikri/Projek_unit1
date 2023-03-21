package main

import (
	"fmt"

	"github.com/wanta-zulfikri/Projek_unit1/app"
	"github.com/wanta-zulfikri/Projek_unit1/db"
	"github.com/wanta-zulfikri/Projek_unit1/db/migration"
	"github.com/wanta-zulfikri/Projek_unit1/repository/produk"
	"github.com/wanta-zulfikri/Projek_unit1/repository/user"
)

func main() {
	DB := db.InitDb()
	migration.Migration()
	InitUser := user.InitUser(DB)
	InitProduk := produk.InitProduk(DB) 
	var choice int
	defer DB.Close()
	defer fmt.Println("Terimakasih telah menggunakan aplikasi kami")
	fmt.Println("Daftar Pilihan :")
	fmt.Println("1.Running Aplikasi")
	fmt.Println("9.Exit")
	fmt.Print("Masukan Pilihan: ")
	fmt.Scanln(&choice)
	for choice != 9 && choice == 1 {
		switch choice {
		case 1:
			App := app.InitApp(InitUser,InitProduk, &choice)
			App.Home()
		}

	}

}
