package app

import (
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/entities"
	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func (app *App) HomePegawai() {
	key := helper.GetUser(app.Session)
	fmt.Printf("Selamat datang ,%s", app.Session[key].Username)
	var choice int
	fmt.Print("\x1bc")
	fmt.Println("\n=============Pilih Menu Dibawah ini======================")
	fmt.Println("1.Tambah Produk")
	fmt.Println("2.Update Produk")
	fmt.Println("3.Tambah Customer")
	fmt.Println("4.Buat Nota")
	fmt.Println("5.Lihat Transaksi")
	fmt.Println("6.Update Profile")
	fmt.Println("7.Logout")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		app.TambahProduk()
	case 2:
		app.UpdateProduk()
	case 3:
		app.TambahCustomer()
	case 4:
		app.CreateTransaction()
	case 5:
		app.ListTransaction()
	case 6:
		app.UpdateAccount()
	case 7:
		app.Logout()
	default:
		app.HomePegawai()

	}
}

func (app *App) UpdateAccount() {
	var username, password string
	key := helper.GetUser(app.Session)
	fmt.Println("=================Form Update Account===============")
	fmt.Print("Masukan Username Baru (Tekan Enter Untuk Skip): ")
	fmt.Scanln(&username)
	if username == "" {
		username = app.Session[key].Username
	} else {
		if username != app.Session[key].Username {
			_, err := app.usersRepo.FindByUsername(username)
			if err == nil {
				fmt.Print("Username Sudah Terdaftar!! Silahkan Ganti!!!")
				time.Sleep(3 * time.Second)
				app.UpdateAccount()
			}
			err1 := app.usersRepo.InsertLog(app.Session[key].Username, username)
			if err1 != nil {
				fmt.Print("Gagal menambahkan log")
				time.Sleep(2 * time.Second)
				app.UpdateAccount()
			}
		}
	}
	fmt.Print("Masukan Password Baru (Tekan Enter Untuk Skip): ")
	fmt.Scanln(&password)
	if password == "" {
		password = app.Session[key].Password
	}
	err := app.usersRepo.Update(&entities.User{Username: username, Password: password}, app.Session[key].Id)
	var choice string
	if err != nil {
		fmt.Print("Gagal mengupdate profile,apakah ingin mencoba lagi? (y/t) ")
		fmt.Scanln(&choice)
		if choice == "y" {
			app.UpdateAccount()
		}
		fmt.Println("Anda akan diarahakan ke halaman utama")
		time.Sleep(time.Second * 3)
		app.HomePegawai()
	}
	if username != app.Session[key].Username {
		newdata, _ := app.usersRepo.FindByUsername(username)
		delete(app.Session, key)
		app.Session[username] = newdata
	}
	fmt.Println("Berhasil mengupdate data,anda akan diarahakan ke halaman utama")
	time.Sleep(time.Second * 3)
	app.HomePegawai()
}
