package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
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

func (cus *App) DeleteCustomer() {
	var choice string
	key := helper.GetUser(cus.Session)
	fmt.Print("\x1bc")
	fmt.Println("==============FORM HAPUS Customer================")
	fmt.Println()
	lenght, _ := cus.CusRepo.GetAll()
	page := helper.CalculatePage(len(lenght))
	datas, _ := cus.CusRepo.GetWithLimit(cus.OffsetContent)
	if len(lenght) == 0 {
		fmt.Println("Data Customer Belum Ada")
		fmt.Print("Apakah anda ingin menambahkan data Customer? (y/t): ")
		fmt.Scanln(&choice)
		if choice == "y" {
			cus.TambahCustomer()
			return
		}
		if cus.Session[key].Role == "admin" {
			cus.HomeAdmin()
			return
		}
		cus.HomePegawai()
		return
	}
	helper.PrintData(datas)
	if page > cus.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter : ")
		fmt.Scanln(&choice)
		if choice == "L" {
			cus.PageContent++
			cus.OffsetContent += config.LimitPage
			cus.DeleteCustomer()
			return
		}
	} else if cus.PageContent != 1 || (cus.PageContent == page && page > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			cus.PageContent--
			cus.OffsetContent -= config.LimitPage
			cus.DeleteCustomer()
			return
		}
	}
	fmt.Print("Silahkan Pilih Customer Yang Ingin Dihapus jika ingin sekaligus gunakan format (ex:1,2,3): ")
	fmt.Scanln(&choice)
	var index int
	if strings.Contains(choice, ",") {
		ids := strings.Split(choice, ",")
		for i, val := range ids {
			if helper.IsNotInt(val) {
				fmt.Print("Wajib Angka!!!,Ingin mencoba lagi? y/t: ")
				fmt.Scanln(&choice)
				if choice == "y" {
					helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
					cus.DeleteCustomer()
				}
				helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
				fmt.Println("Kamu Akan diarahkan ke halaman utama")
				time.Sleep(time.Second * 3)
				if cus.Session[key].Role == "admin" {
					cus.HomeAdmin()
					return
				}
				cus.HomePegawai()
				return
			}
			toint := helper.ConvertStringToInt(val) - 1
			err := cus.CusRepo.Delete(datas[toint].Id)
			if err != nil {
				fmt.Printf("Id %d Tidak terdaftar\n", datas[toint].Id)
				fmt.Printf("Data yang Dihapus sebanyak %d", i+1)
				break
			}
			index++
		}
		if index < 1 {
			helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
			fmt.Println("Masukan Data Yang benar")
			time.Sleep(2 * time.Second)
			cus.DeleteCustomer()
			return
		}
		fmt.Println("Berhasil Mengapus Data")
		fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
		fmt.Scanln(&choice)
		if choice == "y" {
			helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
			cus.DeleteCustomer()
			return
		}
		helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
		fmt.Println("Anda Akan diarahkan ke menu dashboard")
		time.Sleep(time.Second * 2)
		if cus.Session[key].Role == "admin" {
			cus.HomeAdmin()
		}
		cus.HomePegawai()
		return
	}
	if helper.IsNotInt(choice) {
		fmt.Print("Wajib Angka!!!,Ingin mencoba lagi? y/t: ")
		fmt.Scanln(&choice)
		if choice == "y" {
			helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
			cus.DeleteCustomer()
		}
		helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
		fmt.Println("Kamu Akan diarahkan ke halaman utama")
		time.Sleep(time.Second * 3)
		if cus.Session[key].Role == "admin" {
			cus.HomeAdmin()
			return
		}
		cus.HomePegawai()
		return
	}
	err := cus.CusRepo.Delete(datas[helper.ConvertStringToInt(choice)-1].Id)
	if err != nil {
		helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
		fmt.Println("Masukan Data Yang Benar!!!")
		time.Sleep(time.Second * 2)
		cus.DeleteCustomer()
		return
	}
	fmt.Println("Berhasil Mengahapus Data")
	fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
	fmt.Scanln(&choice)
	if choice == "y" {
		helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
		cus.DeleteCustomer()
		return
	}
	helper.ResetValue(&cus.PageContent, &cus.OffsetContent, 1, 0)
	fmt.Println("Anda akan diarahkan ke halaman utama")
	time.Sleep(time.Second * 2)
	if cus.Session[key].Role == "admin" {
		cus.HomeAdmin()
		return
	}
	cus.HomePegawai()
	return
}

func (cus *App) UpdateCustomer() {
	key := helper.GetUser(cus.Session)
	var choice, name, address, phonenum string
	fmt.Print("\x1bc")
	fmt.Println("=============FORM UPDATE CUSTOMER=================")
	fmt.Println()
	lenght, _ := cus.CusRepo.GetAll()
	page := helper.CalculatePage(len(lenght))
	datas, _ := cus.CusRepo.GetWithLimit(cus.OffsetContent)
	if len(lenght) == 0 {
		fmt.Println("Data Customer Belum Ada")
		fmt.Print("Apakah anda ingin menambahkan data pegawai? (y/t): ")
		fmt.Scanln(&choice)
		if choice == "y" {
			cus.UpdateCustomer()
			return
		}
		fmt.Println("Anda Akan Diarahkan Ke Halaman Utama")
		time.Sleep(3 * time.Second)
		if cus.Session[key].Role == "admin" {
			cus.HomeAdmin()
		}
		cus.HomePegawai()
		return
	}
	helper.PrintData(datas)
	if page > cus.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter : ")
		fmt.Scanln(&choice)
		if choice == "L" {
			cus.PageContent++
			cus.OffsetContent += config.LimitPage
			cus.UpdateCustomer()
			return
		}
	} else if cus.PageContent != 1 || (cus.PageContent == page && page > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			cus.PageContent--
			cus.OffsetContent -= config.LimitPage
			cus.UpdateCustomer()
			return
		}
	}
	fmt.Print("Silahkan Pilih Pegawai Yang Ingin Di Update: ")
	fmt.Scanln(&choice)
	if helper.IsNotInt(choice) {
		fmt.Print("Masukan pilihan yang benar!!. Apakah ingin mengulang (y/t)? ")
		fmt.Scanln(&choice)
		if choice == "y" {
			cus.UpdateCustomer()
			return
		}
		fmt.Println("Anda akan diarahkan ke halaman utama")
		time.Sleep(3 * time.Second)
		if cus.Session[key].Role == "admin" {
			cus.HomeAdmin()
		}
		cus.HomePegawai()
		return
	}
	fmt.Print("Masukan Nama Baru (Enter Untuk Skip) : ")
	cus.Scanner.Scan()
	name = cus.Scanner.Text()
	if name == "" {
		name = datas[helper.ConvertStringToInt(choice)-1].Nama
	}
	fmt.Print("Masukan Alamat Baru (Enter Untuk Skip): ")
	cus.Scanner.Scan()
	address = cus.Scanner.Text()
	if address == "" {
		address = datas[helper.ConvertStringToInt(choice)-1].Alamat
	}
	fmt.Print("Masukan No Hp Baru (Enter Untuk Skip): ")
	cus.Scanner.Scan()
	phonenum = cus.Scanner.Text()
	if phonenum == "" {
		address = datas[helper.ConvertStringToInt(choice)-1].NoHp
	}
	err := cus.CusRepo.Update(&entities.Customer{NoHp: phonenum, Nama: name, Alamat: address, Id: datas[helper.ConvertStringToInt(choice)-1].Id})

	if err != nil {
		fmt.Print("Gagal mengupdate user, apakah anda ingin mencoba lagi? (y/t) :  ")
		fmt.Scanln(&choice)
		if choice == "y" {
			cus.UpdateCustomer()
			return
		}
		fmt.Println("Anda akan diarahakan ke menu utama")
		time.Sleep(3 * time.Second)
		if cus.Session[key].Role == "admin" {
			cus.HomeAdmin()
		}
		cus.HomePegawai()
	}
	fmt.Print("Berhasil Mengupdate Data Pegawai, apakah anda ingin mencoba lagi? (y/t) :  ")
	fmt.Scanln(&choice)
	if choice == "y" {
		cus.UpdateCustomer()
		return
	}
	fmt.Println("Anda akan diarahakan ke menu utama")
	if cus.Session[key].Role == "admin" {
		cus.HomeAdmin()
	}
	cus.HomePegawai()
	return

}
