package helper

import (
	"fmt"
	"math"

	"github.com/wanta-zulfikri/Projek_unit1/entities"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
func GetUser(user map[string]*entities.User) string {
	for _, val := range user {
		return val.Username
	}
	return ""
}
func PrintData(datas interface{}) {
	if datas, ok := datas.([]*entities.User); ok {
		fmt.Println("Berikut List Pegawai: ")
		for i, val := range datas {
			fmt.Printf("%d. Nama Pegawai : %s , Password Pegawai: %s\n", i+1, val.Username, val.Password)
		}
	}
}

func CalculatePage(length int) int {
	return int(math.Ceil(float64(length) / float64(5)))
}

func ResetValue(old1 *int, old2 *int, new ...int) {
	*old1 = new[0]
	*old2 = new[1]
}
