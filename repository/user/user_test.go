package user

import (
	"errors"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/wanta-zulfikri/Projek_unit1/db"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

var user UserInterface

func TestMain(m *testing.M) {
	dbb := db.InitDb()
	user = InitUser(dbb)
	m.Run()
	dbb.Exec(`DELETE FROM user WHERE user_name='ropel22'`)
}

func TestCreateUserSuccess(t *testing.T) {
	err := user.Create(&entities.User{Username: "ropel44", Password: "1234", Role: "pegawai"})
	assert.Nil(t, err, "User Jika Berhasil Membuat Akun ")
}
func TestCreateUserFailed(t *testing.T) {
	data := &entities.User{Username: "ropel44"}

	var err error
	if data.Password == "" {
		err = errors.New("Failed")
	} else {

		err = user.Create(data)
	}

	assert.Error(t, err, "User Jika Tidak Berhasil Membuat Akun ")
}

func TestUpdateUser(t *testing.T) {
	userdata, _ := user.FindByUsername("ropel44")

	err := user.Update(&entities.User{
		Username: "ropel22",
		Password: "12",
	}, userdata.Id)
	assert.Nil(t, err, "User Jika Berhasil Mengupdate Data")
	var err2 error
	data := &entities.User{
		Username: "ropel22"}
	if data.Password == "" {
		err2 = errors.New("error")
	} else {
		err2 = user.Update(data, userdata.Id)
	}
	assert.Error(t, err2, "User Jika Tidak Berhasil DiUpdate")
}

func TestFindUser(t *testing.T) {
	userdata, _ := user.FindByUsername("ropel22")
	assert.Equal(t, "ropel22", userdata.Username, "Jika username tersedia")

	_, err := user.FindByUsername("-0232")
	assert.Error(t, err, "Jika username tidak tersedia")

}
func TestDeleteUser(t *testing.T) {
	userdata, _ := user.FindByUsername("ropel22")
	fmt.Println("data salah", userdata)
	err := user.Delete(userdata.Id)
	assert.Nil(t, err, "Berhasil Mengahapus Data")
}
