package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func (trx *App) CreateTransaction() {
	var cusid, produkid, choice string
	key := helper.GetUser(trx.Session)
	fmt.Print("\x1bc")
	fmt.Println("==================FORM BUAT NOTA=======================")
	if !trx.IsNext {
		lenght, _ := trx.CusRepo.GetAll()
		page := helper.CalculatePage(len(lenght))
		datas, _ := trx.CusRepo.GetWithLimit(trx.OffsetContent)
		if len(lenght) == 0 {
			fmt.Println("Data Customer Belum Ada")
			fmt.Print("Apakah anda ingin menambahkan data Customer? (y/t): ")
			fmt.Scanln(&choice)
			if choice == "y" {
				trx.TambahCustomer()
				return
			}
			if trx.Session[key].Role == "admin" {
				trx.HomeAdmin()
				return
			}
			trx.HomePegawai()
			return
		}
		helper.PrintData(datas)
		if page > trx.PageContent {
			fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter : ")
			fmt.Scanln(&choice)
			if choice == "L" {
				trx.PageContent++
				trx.OffsetContent += config.LimitPage
				trx.CreateTransaction()
				return
			}
		} else if trx.PageContent != 1 || (trx.PageContent == page && page > 1) {
			fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
			fmt.Scanln(&choice)
			if choice == "K" {
				trx.PageContent--
				trx.OffsetContent -= config.LimitPage
				trx.CreateTransaction()
				return
			}
		}
		fmt.Print("Silahkan Pilih Customer: ")
		fmt.Scanln(&cusid)
		if helper.IsNotInt(cusid) || cusid == "" {
			fmt.Print("Wajib Angka/Tidak Boleh Kosong!!!,Ingin mencoba lagi? y/t: ")
			fmt.Scanln(&choice)
			if choice == "y" {
				trx.CreateTransaction()
			}
			fmt.Println("Kamu Akan diarahkan ke halaman utama")
			time.Sleep(time.Second * 3)
			if trx.Session[key].Role == "admin" {
				trx.HomeAdmin()
				return
			}
			trx.HomePegawai()
			return
		}
	}
	fmt.Print("\x1bc")
	trx.Cache["cus"] = &entities.Customer{Id: helper.ConvertStringToInt(cusid)}
	trx.IsNext = true
	lenght1, _ := trx.ProdukRepo.GetAll()
	page1 := helper.CalculatePage(len(lenght1))
	datas, _ := trx.ProdukRepo.GetAllProduk(trx.OffsetContent)
	if len(lenght1) == 0 {
		fmt.Println("Data Produk Belum Ada")
		fmt.Print("Apakah anda ingin menambahkan data produk? (y/t): ")
		fmt.Scanln(&choice)
		if choice == "y" {
			trx.TambahProduk()
			return
		}
		trx.IsNext = false
		if trx.Session[key].Role == "admin" {
			trx.HomeAdmin()
		}
		trx.HomePegawai()
		return
	}
	helper.PrintData(datas)
	if page1 > trx.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter :  ")
		fmt.Scanln(&choice)
		if choice == "L" {
			trx.PageContent++
			trx.OffsetContent += config.LimitPage
			trx.UpdateProduk()
			return
		}
	} else if trx.PageContent != 1 || (trx.PageContent == page1 && page1 > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			trx.PageContent++
			trx.OffsetContent -= config.LimitPage
			trx.UpdateProduk()
			return
		}
	}
	fmt.Print("Silahkan Pilih Produk Yang Ingin Ditambahkan,jika ingin banyak produk ex(1,2,3,4): ")
	fmt.Scanln(&produkid)
	if produkid == "" {
		fmt.Print("Wajib Angka/Tidak Boleh Kosong!!!,Ingin mencoba lagi? y/t: ")
		fmt.Scanln(&choice)
		if choice == "y" {
			trx.CreateTransaction()
		}
		trx.IsNext = false
		delete(trx.Cache, "cus")
		fmt.Println("Kamu Akan diarahkan ke halaman utama")
		time.Sleep(time.Second * 3)
		if trx.Session[key].Role == "admin" {
			trx.HomeAdmin()
		}
		trx.HomePegawai()
	}
	var index int
	if strings.Contains(produkid, ",") {
		ids := strings.Split(produkid, ",")
		err, trxid := trx.TrxRepo.Create(&entities.Transaksi{UserId: trx.Session[key].Id, CustomerId: trx.Cache["cus"].Id})
		if err != nil {
			fmt.Print("Gagal membuat nota,Ingin mencoba lagi? y/t: ")
			fmt.Scanln(&choice)
			if choice == "y" {
				trx.CreateTransaction()
			}
			trx.IsNext = false
			delete(trx.Cache, "cus")
			fmt.Println("Kamu Akan diarahkan ke halaman utama")
			time.Sleep(time.Second * 3)
			if trx.Session[key].Role == "admin" {
				trx.HomeAdmin()
				return
			}
			trx.HomePegawai()
			return
		}
		var databerhasil = ""
		for i, val := range ids {
			if helper.IsNotInt(val) {
				fmt.Print(fmt.Sprintf("Data Produk ke %d Bukan Angka,Apakah Ingin Melanjutkan? (y/t)", i+1))
				fmt.Scanln(&choice)
				if choice == "y" {
					repeat := true
					for repeat {
						fmt.Print("\x1bc")
						helper.PrintData(datas)
						fmt.Print("Masukan Pilihan Barang: ")
						fmt.Scanln(&choice)
						if !helper.IsNotInt(choice) && !helper.IfAlreadyTaken(databerhasil, choice) {
							val = choice
							repeat = false
						} else {
							fmt.Print(fmt.Sprintf("\n Data Produk %s Sudah dibuat\n", datas[helper.ConvertStringToInt(choice)-1].Nama_produk))
							time.Sleep(3 * time.Second)
						}
					}
				} else {
					trx.IsNext = false
					delete(trx.Cache, "cus")
					if index > 0 {
						fmt.Printf("Produk Berhasil Ditambahkan Sebanyak %d, Kamu Akan diarahkan ke halaman utama\n", index)
					}
					fmt.Println("Kamu Akan diarahkan ke halaman utama")
					time.Sleep(time.Second * 3)
					if trx.Session[key].Role == "admin" {

						trx.HomeAdmin()
					}
					trx.HomePegawai()
				}
			}
			var qty int
			toint := helper.ConvertStringToInt(val) - 1
			fmt.Print(fmt.Sprintf("Berapa banyak produk %s yang ingin dibeli? : ", datas[toint].Nama_produk))
			fmt.Scanln(&qty)
			if qty == 0 || qty < 0 {
				fmt.Println("kuantiti Tidak Boleh Kosong atau kurang dari 0")
				var repeat = true
				for repeat {
					fmt.Print(fmt.Sprintf("Berapa banyak produk %s yang ingin dibeli? : ", datas[toint].Nama_produk))
					fmt.Scanln(&qty)
					if qty == 0 || qty < 0 {
						repeat = false
					} else {
						fmt.Println("kuantiti Tidak Boleh Kosong atau kurang dari 0")
					}
				}
			}
			newqty := datas[toint].Qty - qty
			if newqty < 0 {
				fmt.Printf("Jumlah stok pada produk %s hanya tersisa %d\n", datas[toint].Nama_produk, datas[toint].Qty)
				var repeat = true
				for repeat {
					fmt.Print(fmt.Sprintf("Berapa banyak produk %s yang ingin dibeli? : ", datas[toint].Nama_produk))
					fmt.Scanln(&qty)
					if qty <= datas[toint].Qty {
						repeat = false
						newqty = datas[toint].Qty - qty
					} else {
						fmt.Printf("Jumlah stok pada produk %s hanya tersisa %d\n", datas[toint].Nama_produk, datas[toint].Qty)
					}
				}
			}
			err := trx.ProdukRepo.UpdateStok(newqty, datas[toint].Id)
			if err != nil {
				fmt.Printf("Id Produk %d Tidak terdaftar\n", datas[toint].Id)
				fmt.Printf("Produk yang Ditambahkan sebanyak %d\n", i+1)
				break
			}
			trx.TrxRepo.InsertItems(&entities.TransaksiItem{Qty: qty, Price: datas[toint].Price, TrxId: trxid, ProdukId: datas[toint].Id})
			index++
			databerhasil += val
		}
		if index < 1 {
			fmt.Println("Masukan produk Yang benar")
			time.Sleep(2 * time.Second)
			trx.CreateTransaction()
		}
		fmt.Println("Berhasil membuat nota")
		fmt.Print("Apakah anda ingin membuat nota lagi? (y/t)")
		fmt.Scanln(&choice)
		if choice == "y" {
			helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
			trx.IsNext = false
			delete(trx.Cache, "cus")
			trx.CreateTransaction()
			return
		}
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		trx.IsNext = false
		delete(trx.Cache, "cus")
		fmt.Println("Anda Akan diarahkan ke menu dashboard")
		time.Sleep(time.Second * 3)
		if trx.Session[key].Role == "admin" {

			trx.HomeAdmin()
		}
		trx.HomePegawai()
		return
	}
	///Single produk
	if helper.IsNotInt(produkid) {
		fmt.Print("Inputan bukan berupa angka,apakah ingin mencoba lagi? ")
		fmt.Scanln(&choice)
		if choice == "y" {
			repeat := true
			for repeat {
				helper.PrintData(datas)
				fmt.Print("Masukan Pilihan Barang: ")
				fmt.Scanln(&produkid)
				if !helper.IsNotInt(produkid) {
					repeat = false
				}
			}
		} else {
			trx.IsNext = false
			delete(trx.Cache, "cus")
			fmt.Println("Kamu Akan diarahkan ke halaman utama")
			time.Sleep(time.Second * 3)
			if trx.Session[key].Role == "admin" {

				trx.HomeAdmin()
			}
			trx.HomePegawai()
		}
	}
	err, trxid := trx.TrxRepo.Create(&entities.Transaksi{UserId: trx.Session[key].Id, CustomerId: trx.Cache["cus"].Id})
	if err != nil {
		fmt.Print("Gagal membuat nota,Ingin mencoba lagi? y/t: ")
		fmt.Scanln(&choice)
		if choice == "y" {
			trx.CreateTransaction()
		} else {
			trx.IsNext = false
			delete(trx.Cache, "cus")
			fmt.Println("Kamu Akan diarahkan ke halaman utama")
			time.Sleep(time.Second * 3)
			if trx.Session[key].Role == "admin" {
				trx.HomeAdmin()
			}
			trx.HomePegawai()
		}
	}
	var qty int
	toint := helper.ConvertStringToInt(produkid) - 1
	fmt.Print(fmt.Sprintf("Berapa banyak produk %s yang ingin dibeli? : ", datas[toint].Nama_produk))
	fmt.Scanln(&qty)
	if qty == 0 || qty < 0 {
		fmt.Println("kuantiti Tidak Boleh Kosong atau kurang dari 0")
		var repeat = true
		for repeat {
			fmt.Print(fmt.Sprintf("Berapa banyak produk %s yang ingin dibeli? : ", datas[toint].Nama_produk))
			fmt.Scanln(&qty)
			if qty == 0 || qty < 0 {
				repeat = false
			} else {
				fmt.Println("kuantiti Tidak Boleh Kosong atau kurang dari 0")
			}
		}
	}
	newqty := datas[toint].Qty - qty
	if newqty < 0 {
		fmt.Printf("Jumlah stok pada produk %s hanya tersisa %d\n", datas[toint].Nama_produk, datas[toint].Qty)
		var repeat = true
		for repeat {
			fmt.Print(fmt.Sprintf("Berapa banyak produk %s yang ingin dibeli? : ", datas[toint].Nama_produk))
			fmt.Scanln(&qty)
			if qty <= datas[toint].Qty {
				repeat = false
				newqty = datas[toint].Qty - qty
			} else {
				fmt.Printf("Jumlah stok pada produk %s hanya tersisa %d\n", datas[toint].Nama_produk, datas[toint].Qty)
			}
		}
	}
	trx.ProdukRepo.UpdateStok(newqty, datas[toint].Id)
	trx.TrxRepo.InsertItems(&entities.TransaksiItem{Qty: qty, Price: datas[toint].Price, TrxId: trxid, ProdukId: datas[toint].Id})
	fmt.Println("Berhasil membuat nota")
	fmt.Print("Apakah anda ingin membuat nota lagi? (y/t)")
	fmt.Scanln(&choice)
	if choice == "y" {
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		trx.IsNext = false
		delete(trx.Cache, "cus")
		trx.CreateTransaction()

	}
	helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
	trx.IsNext = false
	delete(trx.Cache, "cus")
	fmt.Println("Anda Akan diarahkan ke menu dashboard")
	time.Sleep(time.Second * 3)
	if trx.Session[key].Role == "admin" {
		trx.HomeAdmin()
	}
	trx.HomePegawai()
}

func (trx *App) ListTransaction() {
	var choice string
	fmt.Println("\x1bc")
	fmt.Println("========List Transaksi=========")
	key := helper.GetUser(trx.Session)
	var lenght []*entities.Transaksi
	var datas []*entities.Transaksi
	lenght, _ = trx.TrxRepo.GetAllByUid(trx.Session[key].Id)
	datas, _ = trx.TrxRepo.GetWithLimitByUid(trx.Session[key].Id, trx.OffsetContent)
	fmt.Println(lenght, datas)
	if trx.Session[key].Role == "admin" {
		lenght, _ = trx.TrxRepo.GetAll()
		datas, _ = trx.TrxRepo.GetWithLimit(trx.OffsetContent)
	}

	page := helper.CalculatePage(len(lenght))
	if len(lenght) == 0 {
		fmt.Println("Data Transaksi Belum Ada")
		time.Sleep(time.Second * 2)
		if trx.Session[key].Role == "admin" {
			trx.HomeAdmin()
		}
		trx.HomePegawai()
	}
	helper.PrintData(datas, trx.TrxRepo)
	if page > trx.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Ingin Kembali Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "L" {
			trx.PageContent++
			trx.OffsetContent += config.LimitPage
			trx.ListTransaction()

		}
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		if trx.Session[key].Role == "admin" {
			trx.HomeAdmin()
		}
		trx.HomePegawai()
	} else if trx.PageContent != 1 || (trx.PageContent == page && page > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Ingin Kembali Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			trx.PageContent--
			trx.OffsetContent -= config.LimitPage
			trx.ListTransaction()

		}
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		if trx.Session[key].Role == "admin" {
			trx.HomeAdmin()
		}
		trx.HomePegawai()
	}
	fmt.Print("Tekan Enter jika ingin kembali: ")
	fmt.Scanln()
	helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
	if trx.Session[key].Role == "admin" {
		trx.HomeAdmin()
	}
	trx.HomePegawai()

}

func (trx *App) DeleteTransaction() {
	var choice string
	fmt.Print("\x1bc")
	fmt.Println("========Form Delete Transaksi=========")
	lenght, _ := trx.TrxRepo.GetAll()
	datas, _ := trx.TrxRepo.GetWithLimit(trx.OffsetContent)
	page := helper.CalculatePage(len(lenght))
	if len(lenght) == 0 {
		fmt.Println("Data Transaksi Belum Ada")
		time.Sleep(time.Second * 3)
		trx.HomeAdmin()
	}
	helper.PrintData(datas, trx.TrxRepo)
	if page > trx.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Ingin Kembali Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "L" {
			trx.PageContent++
			trx.OffsetContent += config.LimitPage
			trx.DeleteTransaction()

		}
	} else if trx.PageContent != 1 || (trx.PageContent == page && page > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Ingin Kembali Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			trx.PageContent--
			trx.OffsetContent -= config.LimitPage
			trx.DeleteTransaction()
		}
	}
	fmt.Print("Silahkan Pilih transaksi Yang Ingin Dihapus,jika ingin sekaligus gunakan format (ex:1,2,3): ")
	fmt.Scanln(&choice)
	var index int
	if strings.Contains(choice, ",") {
		ids := strings.Split(choice, ",")
		for i, val := range ids {
			if helper.IsNotInt(val) {
				fmt.Print("Wajib Angka!!!,Ingin mencoba lagi? y/t: ")
				fmt.Scanln(&choice)
				if choice == "y" {
					helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
					trx.DeleteTransaction()
				}
				helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
				fmt.Println("Kamu Akan diarahkan ke halaman utama")
				time.Sleep(time.Second * 3)
				trx.HomeAdmin()
			}
			toint := helper.ConvertStringToInt(val) - 1
			err := trx.TrxRepo.Delete(datas[toint].Id)
			if err != nil {
				fmt.Printf("Id %d Tidak terdaftar\n", datas[toint].Id)
				fmt.Printf("Data yang Dihapus sebanyak %d", i+1)
				break
			}
			index++
		}
		if index < 1 {
			helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
			fmt.Println("Masukan Data Yang benar")
			time.Sleep(2 * time.Second)
			trx.DeleteTransaction()

		}
		fmt.Println("Berhasil Mengapus Data")
		fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
		fmt.Scanln(&choice)
		if choice == "y" {
			helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
			trx.DeleteTransaction()
		}
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		fmt.Println("Anda Akan diarahkan ke menu dashboard")
		time.Sleep(time.Second * 2)
		trx.HomeAdmin()
	}
	if helper.IsNotInt(choice) {
		fmt.Print("Wajib Angka!!!,Ingin mencoba lagi? y/t: ")
		fmt.Scanln(&choice)
		if choice == "y" {
			helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
			trx.DeleteTransaction()
		}
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		fmt.Println("Kamu Akan diarahkan ke halaman utama")
		time.Sleep(time.Second * 3)
		trx.HomeAdmin()
	}
	err := trx.TrxRepo.Delete(datas[helper.ConvertStringToInt(choice)-1].Id)
	if err != nil {
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		fmt.Println("Masukan Data Yang Benar!!!")
		time.Sleep(time.Second * 2)
		trx.DeleteTransaction()

	}
	fmt.Println("Berhasil Mengahapus Data")
	fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
	fmt.Scanln(&choice)
	if choice == "y" {
		helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
		trx.DeleteCustomer()

	}
	helper.ResetValue(&trx.PageContent, &trx.OffsetContent, 1, 0)
	fmt.Println("Anda akan diarahkan ke halaman utama")
	time.Sleep(time.Second * 2)
	trx.HomeAdmin()

}
