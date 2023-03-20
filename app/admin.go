package app

import (
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

func (admin *App) HomeAdmin() {
	fmt.Printf("=========Selamat datang admin!!!==============")
	var choice int
	fmt.Println("\n=============Pilih Menu Dibawah ini======================")
	fmt.Println("1.Tambah Pegawai")
	fmt.Println("2.Update Produk")
	fmt.Println("3.Tambah Produk")
	fmt.Println("4.Tambah Produk")
	fmt.Println("5.Tambah Produk")
	fmt.Println("6.Hapus Pegawai")
	fmt.Print("Masukan Pilihan : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		admin.TambahPegawai()
		return
	case 2:
	}

}

func (admin *App) TambahPegawai() {
	var username, password, Repeatlogin string
	fmt.Println("\n=============Form Tambah Pegawai======================")
	fmt.Print("Masukan Username Pegawai: ")
	fmt.Scanln(&username)
	_, err := admin.usersRepo.FindByUsername(username)
	if err != nil {
		fmt.Print("Username sudah terdaftar, apakah ingin mencoba lagi? (y/t) : ")
		fmt.Scan(&Repeatlogin)
		if Repeatlogin == "y" {
			admin.TambahPegawai()
			return
		}
		fmt.Println("Anda akan diarahkan ke halaman utama")
		time.Sleep(time.Second * 3)
		admin.HomeAdmin()
		return
	}
	fmt.Print("Masukan Password: ")
	fmt.Scanln(&password)
	err1 := admin.usersRepo.Create(&entities.User{Username: username, Password: password, Role: "pegawai"})
	if err1 != nil {
		fmt.Println(err1.Error())
		fmt.Println("Anda akan diarahkan ke halaman utama")
	}
	fmt.Println("Berhasil menambahkan pegawai, anda akan di arahkan ke halaman utama")
	time.Sleep(time.Second * 3)
}
