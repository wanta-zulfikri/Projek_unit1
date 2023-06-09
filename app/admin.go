package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func (admin *App) HomeAdmin() {
	fmt.Print("\x1bc")
	fmt.Printf("=========Selamat datang admin!!!==============")
	var choice int
	fmt.Println("\n=============Pilih Menu Dibawah ini======================")
	fmt.Println("1.Tambah Pegawai")
	fmt.Println("2.Update Pegawai")
	fmt.Println("3.Hapus Pegawai")
	fmt.Println("4.Tambah Pelanggan")
	fmt.Println("5.Hapus Pelanggan")
	fmt.Println("6.Tambah Produk")
	fmt.Println("7.Update Produk")
	fmt.Println("8.Hapus Produk")
	fmt.Println("9.Edit Pelanggan")
	fmt.Println("10.Buat Nota")
	fmt.Println("11.Lihat Transaksi")
	fmt.Println("12.Lihat Log Pegawai")
	fmt.Println("13.Delete Transaksi")
	fmt.Println("14.Logout")
	fmt.Print("Masukan Pilihan : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		admin.TambahPegawai()

	case 2:
		admin.UpdatePegawai()

	case 3:
		admin.HapusPegawai()

	case 4:
		admin.TambahCustomer()

	case 5:
		admin.DeleteCustomer()

	case 6:
		admin.TambahProduk()

	case 7:
		admin.UpdateProduk()

	case 8:
		admin.HapusProduk()

	case 9:
		admin.UpdateCustomer()
	case 10:
		admin.CreateTransaction()
	case 11:
		admin.ListTransaction()
	case 12:
		admin.LogAccountPegawai()
	case 13:
		admin.DeleteTransaction()
	case 14:
		admin.Logout()
	default:
		admin.HomeAdmin()
	}

}

func (admin *App) TambahPegawai() {
	var username, password, Repeatlogin string
	fmt.Println("\x1bc")
	fmt.Println("=============Form Tambah Pegawai======================")
	fmt.Print("Masukan Username Pegawai: ")
	fmt.Scanln(&username)
	_, err := admin.usersRepo.FindByUsername(username)
	if err == nil {
		fmt.Print("Username sudah terdaftar, apakah ingin mencoba lagi? (y/t) : ")
		fmt.Scanln(&Repeatlogin)
		if Repeatlogin == "y" {
			admin.TambahPegawai()

		}
		fmt.Println("Anda akan diarahkan ke halaman utama")
		time.Sleep(time.Second * 3)
		admin.HomeAdmin()

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
	admin.HomeAdmin()

}
func (admin *App) HapusPegawai() {
	var choice string
	fmt.Print("\x1bc")
	fmt.Println("==============FORM HAPUS PEGAWI================")
	fmt.Println()
	lenght, _ := admin.usersRepo.GetAllByRole("pegawai")
	page := helper.CalculatePage(len(lenght))
	datas, _ := admin.usersRepo.GetAllByRoleLimit("pegawai", admin.OffsetContent)
	if len(lenght) == 0 {
		fmt.Println("Data Pergawai Belum Ada")
		fmt.Print("Apakah anda ingin menambahkan data pegawai? (y/t): ")
		fmt.Scanln(&choice)
		if choice == "y" {
			admin.TambahPegawai()

		}
		admin.HomeAdmin()

	}
	helper.PrintData(datas)
	if page > admin.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter : ")
		fmt.Scanln(&choice)
		if choice == "L" {
			admin.PageContent++
			admin.OffsetContent += config.LimitPage
			admin.HapusPegawai()

		}
	} else if admin.PageContent != 1 || (admin.PageContent == page && page > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			admin.PageContent--
			admin.OffsetContent -= config.LimitPage
			admin.HapusPegawai()

		}
	}
	fmt.Print("Silahkan Pilih Pegawai Yang Ingin Dihapus jika ingin sekaligus gunakan format (ex:1,2,3): ")
	fmt.Scanln(&choice)
	var index int
	if strings.Contains(choice, ",") {
		ids := strings.Split(choice, ",")
		for i, val := range ids {
			if helper.IsNotInt(val) {
				fmt.Print("Wajib Angka!!!,Ingin mencoba lagi? y/t: ")
				fmt.Scanln(&choice)
				if choice == "y" {
					admin.HapusPegawai()
				}
				fmt.Println("Kamu Akan diarahkan ke halaman utama")
				time.Sleep(time.Second * 3)
				admin.HapusPegawai()
			}
			toint := helper.ConvertStringToInt(val) - 1
			err := admin.usersRepo.Delete(datas[toint].Id)
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
			admin.HapusPegawai()

		}
		fmt.Println("Berhasil Mengapus Data")
		fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
		fmt.Scanln(&choice)
		if choice == "y" {
			helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
			admin.HapusPegawai()

		}
		helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
		fmt.Println("Anda Akan diarahkan ke menu dashboard")
		time.Sleep(time.Second * 2)
		admin.HomeAdmin()

	}
	err := admin.usersRepo.Delete(datas[helper.ConvertStringToInt(choice)-1].Id)
	helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
	if err != nil {
		fmt.Println("Masukan Data Yang Benar!!!")
		time.Sleep(time.Second * 2)
		admin.HapusPegawai()

	}
	helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
	fmt.Println("Berhasil Mengahapus Data")
	fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
	fmt.Scanln(&choice)
	if choice == "y" {
		helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
		admin.HapusPegawai()

	}
	helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
	fmt.Println("Anda akan diarahkan ke halaman utama")
	time.Sleep(time.Second * 2)
	admin.HomeAdmin()
}

func (admin *App) LogAccountPegawai() {
	fmt.Print("\x1bc")
	var choice string
	key := helper.GetUser(admin.Session)
	fmt.Println("=====================Log Akun Pegawai=====================")
	length, _ := admin.usersRepo.GetLog()
	page := helper.CalculatePage(len(length))
	if len(length) == 0 {
		fmt.Println("Belum ada data log pegawai")
		time.Sleep(time.Second * 3)
		if admin.Session[key].Role == "admin" {
			admin.HomeAdmin()
		}
		admin.HomePegawai()
	}
	logs, _ := admin.usersRepo.GetLogWithLimit(admin.OffsetContent)
	helper.PrintData(logs)
	if page > admin.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Ingin Kembali Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "L" {
			admin.PageContent++
			admin.OffsetContent += config.LimitPage
			admin.LogAccountPegawai()

		}
		helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
		if admin.Session[key].Role == "admin" {
			admin.HomeAdmin()
		}
		admin.HomePegawai()
	} else if admin.PageContent != 1 || (admin.PageContent == page && page > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Ingin Kembali Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			admin.PageContent--
			admin.OffsetContent -= config.LimitPage
			admin.LogAccountPegawai()

		}
		helper.ResetValue(&admin.PageContent, &admin.OffsetContent, 1, 0)
		if admin.Session[key].Role == "admin" {
			admin.HomeAdmin()
		}
		admin.HomePegawai()
	}

}
func (admin *App) UpdatePegawai() {
	var choice, username, password string
	fmt.Print("\x1bc")
	fmt.Println("=============FORM UPDATE PEGAWAI=================")
	fmt.Println()
	lenght, _ := admin.usersRepo.GetAllByRole("pegawai")
	page := helper.CalculatePage(len(lenght))
	datas, _ := admin.usersRepo.GetAllByRoleLimit("pegawai", admin.OffsetContent)
	if len(lenght) == 0 {
		fmt.Println("Data Pergawai Belum Ada")
		fmt.Print("Apakah anda ingin menambahkan data pegawai? (y/t): ")
		fmt.Scanln(&choice)
		if choice == "y" {
			admin.TambahPegawai()

		}
		admin.HomeAdmin()

	}
	helper.PrintData(datas)
	if page > admin.PageContent {
		fmt.Print("Tekan L Untuk Page Selanjutnya Dan Jika Tidak Tekan Enter : ")
		fmt.Scanln(&choice)
		if choice == "L" {
			admin.PageContent++
			admin.OffsetContent += config.LimitPage
			admin.UpdatePegawai()

		}
	} else if admin.PageContent != 1 || (admin.PageContent == page && page > 1) {
		fmt.Print("Tekan K Untuk Page Sebelumnya Dan Jika Tidak Tekan Enter: ")
		fmt.Scanln(&choice)
		if choice == "K" {
			admin.PageContent--
			admin.OffsetContent -= config.LimitPage
			admin.UpdatePegawai()

		}
	}
	fmt.Print("Silahkan Pilih Pegawai Yang Ingin Di Update: ")
	fmt.Scanln(&choice)
	if helper.IsNotInt(choice) {
		fmt.Print("Masukan pilihan yang benar!!. Apakah ingin mengulang (y/t)? ")
		fmt.Scanln(&choice)
		if choice == "y" {
			admin.UpdatePegawai()

		}
		fmt.Println("Anda akan diarahkan ke halaman utama")
		time.Sleep(3 * time.Second)
		admin.HomeAdmin()

	}
	fmt.Print("Masukan Username Baru (Enter Untuk Skip) : ")
	fmt.Scanln(&username)
	if username == "" {
		username = datas[helper.ConvertStringToInt(choice)-1].Username
	} else {
		data, err := admin.usersRepo.FindByUsername(username)
		if err == nil && data.Username != datas[helper.ConvertStringToInt(choice)-1].Username {
			fmt.Print("Username sudah tersedia,apakah ingin mengulang (y/t)? ")
			fmt.Scanln(&choice)
			if choice == "y" {
				admin.UpdatePegawai()

			}
			fmt.Println("Anda akan diarahkan ke halaman utama")
			time.Sleep(3 * time.Second)
			admin.HomeAdmin()

		}
	}
	fmt.Print("Masukan Password Baru Anda (Enter Untuk Skip): ")
	fmt.Scanln(&password)
	if password == "" {
		password = datas[helper.ConvertStringToInt(choice)-1].Password
	}
	err := admin.usersRepo.Update(&entities.User{Username: username, Password: password}, datas[helper.ConvertStringToInt(choice)-1].Id)

	if err != nil {
		fmt.Print("Gagal mengupdate user, apakah anda ingin mencoba lagi? (y/t) :  ")
		fmt.Scanln(&choice)
		if choice == "y" {
			admin.UpdatePegawai()

		}
		fmt.Println("Anda akan diarahakan ke menu utama")
		time.Sleep(3 * time.Second)
		admin.HomeAdmin()

	}
	fmt.Print("Berhasil Mengupdate Data Pegawai, apakah anda ingin mencoba lagi? (y/t) :  ")
	fmt.Scanln(&choice)
	if choice == "y" {
		admin.UpdatePegawai()

	}
	fmt.Println("Anda akan diarahakan ke menu utama")
	time.Sleep(3 * time.Second)
	admin.HomeAdmin()

}
