package app

import (
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/entities"
	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func (app *App) InitProduk() {
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
func (app *App) TambahProduk() {
	var Produk string
	key := helper.GetUser(app.Session)
	var choice string
	fmt.Println("\n=============Form Produk======================")
	fmt.Print("masukkan nama produk:  ")
    app.Scanner.Scan()
    Produk = app.Scanner.Text()
	 _, err1:=app.ProdukRepo.Getbynama(Produk)
	if err1 == nil {
		fmt.Print("nama produk sudah terdaftar, apakah anda ingin mencoba lagi ? y/t ")
	    fmt.Scanln(&choice)
		if choice == "y" {
			app.TambahProduk()
			return 
		}
	
	if app.Session[key].Role=="admin"{
		fmt.Println("anda akan diarahkan kehalaman utama admin")
		app.HomeAdmin()
		time.Sleep(time.Second*3)
		return
	}
	fmt.Println("anda akan di arahkan halaman utama pegawai")
	time.Sleep(time.Second*3)
	app.HomePegawai()
	return
}
	err := app.ProdukRepo.Tambahproduk(&entities.Produk{Nama_produk: Produk}, app.Session[key].Id)
	if err != nil { 
		fmt.Print(err.Error()," apakah anda ingin mencoba lagi ? y/t ")
		fmt.Scanln(&choice)
		if choice == "y" {
			app.TambahProduk()
			return 
		} 
		if app.Session[key].Role=="admin"{
			fmt.Println("anda akan diarahkan kehalaman utama admin")
			app.HomeAdmin()
			time.Sleep(time.Second*3)
			return
		}
		fmt.Println("anda akan di arahkan halaman utama pegawai")
		time.Sleep(time.Second*3)
		app.HomePegawai()
		return 

    } 
	fmt.Println("produk berhasil ditambahkan")
	fmt.Print(" apakah anda ingin mencoba lagi ? y/t ")
		fmt.Scanln(&choice)
		if choice == "y"  {
			app.TambahProduk()
			return 
		} 
		if app.Session[key].Role=="admin"{
			fmt.Println("berhasil menambahkan produk")
			fmt.Println("anda akan diarahkan kehalaman utama admin")
			time.Sleep(time.Second*3)
			app.HomeAdmin()
			return
		}
		fmt.Println("berhasil menambahkan produk")
		fmt.Println("anda akan di arahkan halaman utama pegawai")
		time.Sleep(time.Second*3)
		app.HomePegawai()
		

}




	// func (app *App) UpdateProduk() {  

// }

