package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
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
	var Qty int
	key := helper.GetUser(app.Session)
	var choice string
	fmt.Println("\n=============Form Produk======================")
	fmt.Print("masukkan nama produk:  ")
	app.Scanner.Scan()
	Produk = app.Scanner.Text()
	fmt.Print("masukkan jumlah Qty produk: ")
	fmt.Scanln(&Qty)
	_, err1 := app.ProdukRepo.Getbynama(Produk)
	if err1 == nil {
		fmt.Print("nama produk sudah terdaftar, apakah anda ingin mencoba lagi ? y/t ")
		fmt.Scanln(&choice)
		if choice == "y" {
			app.TambahProduk()
			return
		}

		if app.Session[key].Role == "admin" {
			fmt.Println("anda akan diarahkan kehalaman utama admin")
			app.HomeAdmin()
			time.Sleep(time.Second * 3)
			return
		}
		fmt.Println("anda akan di arahkan halaman utama pegawai")
		time.Sleep(time.Second * 3)
		app.HomePegawai()
		return
	}
	err := app.ProdukRepo.Tambahproduk(&entities.Produk{Nama_produk: Produk, Qty: Qty}, app.Session[key].Id)
	if err != nil {
		fmt.Print(err.Error(), " apakah anda ingin mencoba lagi ? y/t ")
		fmt.Scanln(&choice)
		if choice == "y" {
			app.TambahProduk()
			return
		}
		if app.Session[key].Role == "admin" {
			fmt.Println("anda akan diarahkan kehalaman utama admin")
			app.HomeAdmin()
			time.Sleep(time.Second * 3)
			return
		}
		fmt.Println("anda akan di arahkan halaman utama pegawai")
		time.Sleep(time.Second * 3)
		app.HomePegawai()
		return

	}
	fmt.Println("produk berhasil ditambahkan")
	fmt.Print(" apakah anda ingin mencoba lagi ? y/t ")
	fmt.Scanln(&choice)
	if choice == "y" {
		app.TambahProduk()
		return
	}
	if app.Session[key].Role == "admin" {
		fmt.Println("berhasil menambahkan produk")
		fmt.Println("anda akan diarahkan kehalaman utama admin")
		time.Sleep(time.Second * 3)
		app.HomeAdmin()
		return
	}
	fmt.Println("berhasil menambahkan produk")
	fmt.Println("anda akan di arahkan halaman utama pegawai")
	time.Sleep(time.Second * 3)
	app.HomePegawai()

}

func (app *App) UpdateProduk() {
	var choice, namaproduk string
	key := helper.GetUser(app.Session)
	var qty int
	fmt.Print("\x1bc")
	fmt.Println("\n=============Update Produk======================")
	fmt.Println()
	lenght, _ := app.ProdukRepo.GetAll()
	page1 := helper.CalculatePage(len(lenght))
	datas, _ := app.ProdukRepo.GetAllProduk(app.OffsetContent)
	if len(lenght) == 0 {
		fmt.Println("Data Produk Belum Ada")
		fmt.Print("Apakah anda ingin menambahkan data produk? (y/t): ")
		fmt.Scanln(&choice)
		if choice == "y" {
			app.TambahProduk()
			return
		}
		if app.Session[key].Role == "admin" {

			app.HomeAdmin()
		}
		app.HomePegawai()
		return
	}
	helper.PrintData(datas)
	if page1 > app.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter :  ")
		fmt.Scanln(&choice)
		if choice == "L" {
			app.PageContent++
			app.OffsetContent += config.LimitPage
			app.UpdateProduk()
			return
		}
	} else if app.PageContent != 1 || (app.PageContent == page1 && page1 > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			app.PageContent++
			app.OffsetContent -= config.LimitPage
			app.UpdateProduk()
			return
		}
	}
	fmt.Print("Silahkan Pilih Produk Yang Ingin Di Update: ")
	fmt.Scanln(&choice)
	fmt.Print("Masukan Produk Baru (Enter Untuk Skp) : ")
	app.Scanner.Scan()
	namaproduk = app.Scanner.Text()
	if namaproduk == "" {
		namaproduk = datas[helper.ConvertStringToInt(choice)-1].Nama_produk
	} else {
		_, err := app.ProdukRepo.Getbynama(namaproduk)
		if err == nil {
			fmt.Print("Produk sudah tersedia, apakah ingin mengulang (y/t)? ")
			fmt.Scanln(&choice)
			if choice == "y" {
				app.UpdateProduk()
				return
			}
			fmt.Println("Anda akan diarahkan ke halaman utama")
			time.Sleep(3 * time.Second)
			app.HomeAdmin()
			return
		}
	}
	fmt.Print("Masukkan Qty Baru Anda (Enter Untuk Skip): ")
	fmt.Scanln(&qty)
	if qty == 0 {
		qty = datas[helper.ConvertStringToInt(choice)-1].Qty
	}
	err := app.ProdukRepo.UpdateProduk(&entities.Produk{Nama_produk: namaproduk, Qty: qty, Id: datas[helper.ConvertStringToInt(choice)-1].Id}, app.Session[key].Username)

	if err != nil {
		fmt.Print("Gagal mengupdate produk, apakah ingin mencoba lagi? (y/t) ")
		fmt.Scanln(&choice)
		if choice == "y" {
			app.UpdateProduk()
			return
		}
		return
	}
	fmt.Println("Anda akan diarahkan ke menu utama")
	fmt.Print("Berhasil Mengupdate Data Produk, apakah anda ingin mencoba lagi? (y/t): ")
	fmt.Scanln(&choice)
	if choice == "y" {
		app.UpdateProduk()
	}
	time.Sleep(3 * time.Second)
	if app.Session[key].Role == "admin" {

		app.HomeAdmin()
	}
	app.HomePegawai()
	return

}

func (app *App) HapusProduk() {
	var choice string
	key := helper.GetUser(app.Session)
	fmt.Print("\x1bc")
	fmt.Println("\n=============Hapus Produk======================")
	fmt.Println()
	lenght, _ := app.ProdukRepo.GetAll()
	page1 := helper.CalculatePage(len(lenght))
	datas, _ := app.ProdukRepo.GetAllProduk(app.OffsetContent)
	if len(lenght) == 0 {
		fmt.Println("Hapus Data Anda")
		fmt.Print("Apakah anda ingin menghapus data produk? (y/t): ")
		fmt.Scanln(&choice)
		if choice == "y" {
			app.HapusProduk()
			return
		}
		if app.Session[key].Role == "admin" {

			app.HomeAdmin()
		}
		app.HomePegawai()
		return
	}
	helper.PrintData(datas)
	if page1 > app.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter :  ")
		fmt.Scanln(&choice)
		if choice == "L" {
			app.PageContent++
			app.OffsetContent += config.LimitPage
			app.HapusProduk()
			return
		}
	} else if app.PageContent != 1 || (app.PageContent == page1 && page1 > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			app.PageContent++
			app.OffsetContent -= config.LimitPage
			app.HapusProduk()
			return
		}
	}
	fmt.Print("Silahkan Pilih Produk Yang Ingin Di Hapus: ")
	fmt.Scanln(&choice)
	var index int
	if strings.Contains(choice, ",") {
		ids := strings.Split(choice, ",")
		for i, val := range ids {
			if helper.IsNotInt(val) {
				fmt.Print("Wajib Angka!!!,Ingin mencoba lagi? y/t: ")
				fmt.Scanln(&choice)
				if choice == "y" {
					app.HapusProduk()
				}
				fmt.Println("Kamu Akan diarahkan ke halaman utama")
				time.Sleep(time.Second * 3)
				app.HapusProduk()
			}
			toint := helper.ConvertStringToInt(val) - 1
			err := app.ProdukRepo.Delete(datas[toint].Id)
			if err != nil {
				fmt.Printf("Id %d Tidak terdaftar\n", datas[toint].Id)
				fmt.Printf("Data yang Dihapus sebanyak %d", i+1)
				break
			}
			index++
		}
		if index < 1 {
			fmt.Println("Masukan Data Yang benar")
			time.Sleep(2 * time.Second)
			app.HapusProduk()
			return
		}
		fmt.Println("Berhasil Mengapus Data")
		fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
		fmt.Scanln(&choice)
		if choice == "y" {
			helper.ResetValue(&app.PageContent, &app.OffsetContent, 1, 0)
			app.HapusProduk()
			return
		}
		helper.ResetValue(&app.PageContent, &app.OffsetContent, 1, 0)
		fmt.Println("Anda Akan diarahkan ke menu dashboard")
		time.Sleep(time.Second * 2)
		if app.Session[key].Role == "admin" {

			app.HomeAdmin()
		}
		app.HomePegawai()
		return
	}

	err := app.ProdukRepo.Delete(datas[helper.ConvertStringToInt(choice)-1].Id)
	helper.ResetValue(&app.PageContent, &app.OffsetContent, 1, 0)
	if err != nil {
		fmt.Println("Masukan Data Yang Benar!!!")
		time.Sleep(time.Second * 2)
		app.HapusProduk()
		return
	}

	helper.ResetValue(&app.PageContent, &app.OffsetContent, 1, 0)
	fmt.Println("Berhasil Mengahapus Data")
	fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
	fmt.Scanln(&choice)
	if choice == "y" {
		helper.ResetValue(&app.PageContent, &app.OffsetContent, 1, 0)
		app.HapusProduk()
		return
	}
	helper.ResetValue(&app.PageContent, &app.OffsetContent, 1, 0)
	fmt.Println("Anda akan diarahkan ke halaman utama")
	time.Sleep(time.Second * 2)
	if app.Session[key].Role == "admin" {

		app.HomeAdmin()
	}
	app.HomePegawai()

}
