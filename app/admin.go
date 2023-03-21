package app

import "fmt"

func (admin *App) HomeAdmin() {
	fmt.Printf("=========Selamat datang admin!!!==============")
	var choice int
	fmt.Println("\n=============Pilih Menu Dibawah ini======================")
	fmt.Println("1.Tambah Pegawai")
	fmt.Println("2.Hapus Pegawai")
	fmt.Println("3.Tambah Produk")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
	case 2:
	case 3:
	admin.TambahProduk()
	return
	}

}
