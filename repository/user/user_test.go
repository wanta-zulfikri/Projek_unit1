package user

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanta-zulfikri/Projek_unit1/db"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

var Db *sql.DB
var user User
var id int

func TestMain(m *testing.M) {
	dbb := db.InitDb()
	user.db = dbb
	m.Run()
	dbb.Exec("DELETE FROM user WHERE user_name=ropel44")
}

var DB *sql.DB

func TestCreateUserSuccess(t *testing.T) {
	err := user.Create(&entities.User{Username: "ropel44", Password: "1234", Role: "pegawai"})
	assert.Nil(t, err, "User Jika Berhasil Membuat Akun ")
	err2 := user.Create(&entities.User{Username: "ropel44", Password: "1234"})
	assert.Error(t, err2, "User Jika Berhasil Membuat Akun ")
}

func TestUpdateUser(t *testing.T) {
	userdata, _ := user.FindByUsername("ropel44")

	err := user.Update(&entities.User{
		Username: "ropel22",
		Password: "12",
	}, userdata.Id)
	assert.Nil(t, err, "User Jika Berhasil Mengupdate Data")
	err2 := user.Update(&entities.User{
		Username: "ropel22",
	}, userdata.Id)
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
	err := user.Delete(userdata.Id)
	assert.Nil(t, err, "Berhasil Mengahapus Data")
}
