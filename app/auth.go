package app

import (
	"fmt"
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
			return
		}
		app.Home()
		return

	}
	fmt.Printf("Password : ")
	fmt.Scanln(&password)
	if data.Password != password {
		fmt.Println("Password Anda Salah")
		fmt.Print("Login Lagi? (y/t): ")
		fmt.Scanln(&Repeatlogin)
		if Repeatlogin == "y" {
			app.Login()
			return
		}
		app.Home()
		return

	}
	app.Session[data.Username] = data
	if data.Role == "admin" {
		app.HomeAdmin()
		return

	}
	app.HomePegawai()
	return
}
