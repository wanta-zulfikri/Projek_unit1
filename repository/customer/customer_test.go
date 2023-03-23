package customer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanta-zulfikri/Projek_unit1/db"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

var cus CustomerInterface

func TestMain(m *testing.M) {
	dbb := db.InitDb()
	cus = InitCustomer(dbb)
	m.Run()
	dbb.Exec(`DELETE FROM user WHERE user_name='ropel22'`)
}

func TestCreateCustomer(t *testing.T) {
	data := &entities.Customer{NoHp: "081232132", Alamat: "perum aaa", Nama: "satrio"}
	err := cus.Create(data)
	assert.Nil(t, err, "Jika berhasil membuat customer")
	var err1 error
	data2 := &entities.Customer{NoHp: "081232132"}
	if data2.Alamat == "" {
		err1 = errors.New("Data Error")
	} else {
		err1 = cus.Create(data2)
	}
	assert.Error(t, err1, "Jika gagal menambahkan customer")
}
func TestUpdateCustomer(t *testing.T) {
	data, _ := cus.FindByPhone("081232132")
	cust := &entities.Customer{NoHp: "081232132", Alamat: "wwwwww", Nama: "satrio", Id: data.Id}
	err := cus.Update(cust)
	assert.Nil(t, err, "Jika berhasil mengupdate customer")
	err2 := cus.Update(&entities.Customer{NoHp: "081232132"})
	assert.Error(t, err2, "Jika gagal mengupdate customer")
}

func TestDeleteCustomer(t *testing.T) {
	data, _ := cus.FindByPhone("081232132")
	err := cus.Delete(data.Id)
	assert.Nil(t, err, "Jika berhasil delete customer")
	err2 := cus.Delete(9999)
	assert.Error(t, err2, "Jika tidak berhasil delete customer")

}

func TestGetCustomer(t *testing.T) {
	data, _ := cus.FindByPhone("081232132")
	assert.Equal(t, "081232132", data.NoHp, "Jika terdapat data")
	_, err := cus.FindByPhone("11111111111111111111111")
	assert.Error(t, err, "Jika tidak terdapat data")
}
