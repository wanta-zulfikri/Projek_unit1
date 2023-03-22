package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wanta-zulfikri/Projek_unit1/entities"
	"github.com/wanta-zulfikri/Projek_unit1/repository/customer"
	"github.com/wanta-zulfikri/Projek_unit1/repository/produk"
	"github.com/wanta-zulfikri/Projek_unit1/repository/transaksi"
	"github.com/wanta-zulfikri/Projek_unit1/repository/user"
)

type App struct {
	usersRepo     user.UserInterface
	ProdukRepo    produk.ProdukInterface
	CusRepo       customer.CustomerInterface
	TrxRepo       transaksi.TransaksiInterface
	IsNext        bool
	Session       map[string]*entities.User
	Scanner       *bufio.Scanner
	MainChoice    *int
	OffsetContent int
	PageContent   int
	Cache         map[string]*entities.Customer
}

func InitApp(UserRepo user.UserInterface, ProdukRepo produk.ProdukInterface, CusRepo customer.CustomerInterface, TrxRepo transaksi.TransaksiInterface, MainChoice *int) *App {
	return &App{
		usersRepo:     UserRepo,
		ProdukRepo:    ProdukRepo,
		CusRepo:       CusRepo,
		TrxRepo:       TrxRepo,
		Session:       make(map[string]*entities.User, 0),
		Scanner:       bufio.NewScanner(os.Stdin),
		MainChoice:    MainChoice,
		OffsetContent: 0,
		PageContent:   1,
		IsNext:        false,
		Cache:         make(map[string]*entities.Customer, 0),
	}
}

func (app *App) Home() {
	fmt.Print("\x1bc")
	var choice int
	fmt.Println("Silahkan Pilih Menu Dibawah Ini : ")
	fmt.Println("1.Login")
	fmt.Println("3.Kembali")
	fmt.Print("Masukan pilihan : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		app.Login()
		return

	default:
		*app.MainChoice = choice
		return
	}

}
