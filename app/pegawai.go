package app

import (
	"fmt"

	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func (app *App) HomePegawai() {
	key := helper.GetUser(app.Session)
	fmt.Printf("Selamat datang ,%s", app.Session[key].Username)
	var choice int
	fmt.Println("\n=============Pilih Menu Dibawah ini======================")
	fmt.Println("1.Tambah Produk")
	fmt.Println("2.Update Produk")
	fmt.Println("3.Buat Nota")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
	case 2:
	case 3:
	}
}
