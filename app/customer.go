package app

import (
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/entities"
	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func (cus *App) TambahCustomer() {
	var name, address, phonenum string
	key := helper.GetUser(cus.Session)
	fmt.Println("\x1bc")
	fmt.Println("=============Form Tambah Customer======================")
	fmt.Print("Masukan Nama Customer Pegawai: ")
	cus.Scanner.Scan()
	name = cus.Scanner.Text()
	if helper.IsEmpty(name) {
		fmt.Println("Nama tidak boleh kosong")
		time.Sleep(time.Second * 2)
		cus.TambahCustomer()
	}
	fmt.Print("Masukan Alamat Customer : ")
	cus.Scanner.Scan()
	address = cus.Scanner.Text()
	if helper.IsEmpty(name) {
		fmt.Println("Alamat tidak boleh kosong")
		time.Sleep(time.Second * 2)
		cus.TambahCustomer()
	}
	fmt.Print("Masukan Nomer Handphone: ")
	fmt.Scanln(&phonenum)
	_, err := cus.CusRepo.FindByPhone(phonenum)
	if err == nil {
		fmt.Print("No Hp Sudah Terdaftar Silahkan Ganti")
		time.Sleep(time.Second * 2)
	}

	err1 := cus.CusRepo.Create(&entities.Customer{NoHp: phonenum, Nama: name, Alamat: address})
	if cus.Session[key].Role == "admin" {
		if err1 != nil {
			fmt.Println(err1.Error())
			fmt.Println("Anda akan diarahkan ke halaman utama")
			time.Sleep(time.Second * 3)
			cus.HomeAdmin()
			return
		}
		fmt.Println("Berhasil menambahkan customer, anda akan di arahkan ke halaman utama")
		time.Sleep(time.Second * 3)
		cus.HomeAdmin()
		return
	}

	if err1 != nil {
		fmt.Println(err1.Error())
		fmt.Println("Anda akan diarahkan ke halaman utama")
		time.Sleep(time.Second * 3)
		cus.HomePegawai()
	}
	fmt.Println("Berhasil menambahkan customer, anda akan di arahkan ke halaman utama")
	time.Sleep(time.Second * 3)
	cus.HomePegawai()
	return

}
