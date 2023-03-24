package app

import (
	"fmt"

	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func (app *App) Login() {
	fmt.Print("\x1bc")
	var username, password, Repeatlogin string
	fmt.Println("\n====================================================Login Form================================================================")
	fmt.Printf("Username : ")
	fmt.Scanln(&username)
	data, err := app.usersRepo.FindByUsername(username)
	if err != nil {
		fmt.Println("Username Tidak Ditemukan")
		fmt.Print("Login Again?(y/t): ")
		fmt.Scanln(&Repeatlogin)
		if Repeatlogin == "y" {
			app.Login()
		}
		app.Home()
	}
	fmt.Printf("Password : ")
	fmt.Scanln(&password)
	if data.Password != password {
		fmt.Println("Password Anda Salah")
		fmt.Print("Login Lagi? (y/t): ")
		fmt.Scanln(&Repeatlogin)
		if Repeatlogin == "y" {
			app.Login()

		}
		app.Home()

	}
	app.Session[data.Username] = data
	if data.Role == "admin" {
		app.HomeAdmin()

	}
	app.HomePegawai()

}

func (app *App) Logout() {
	key := helper.GetUser(app.Session)
	delete(app.Session, key)
	app.Home()
}
