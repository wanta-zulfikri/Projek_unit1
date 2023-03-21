package helper

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/wanta-zulfikri/Projek_unit1/config"
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
	if datas, ok := datas.([]*entities.Customer); ok {
		fmt.Println("Berikut List Customer: ")
		for i, val := range datas {
			fmt.Printf("%d. Nama Customer : %s\nDetail :\nAlamat: %s\nNo HP : %s\n\n", i+1, val.Nama, val.Alamat, val.NoHp)
		}
	}
	if datas, ok := datas.([]*entities.Produk); ok {
		fmt.Println("Berikut List Produk: ")
		for i, val := range datas {
			if !val.Nama_Pengganti.Valid {
				val.Nama_Pengganti.String = "Belum Ada"
			}
			fmt.Printf("%d. Nama Produk : %s\nQty :%d\nDibuat oleh: %s\nTerakhir Update:%s\n\n", i+1, val.Nama_produk, val.Qty, val.Nama_pembuat, val.Nama_Pengganti.String)
		}
	}
}

func CalculatePage(length int) int {
	return int(math.Ceil(float64(length) / float64(config.LimitPage)))
}

func ResetValue(old1 *int, old2 *int, new ...int) {
	*old1 = new[0]
	*old2 = new[1]
}
func ConvertStringToInt(stringg string) int {
	res, _ := strconv.Atoi(stringg)
	return res
}
func IsNotInt(stringg string) bool {
	regex, _ := regexp.Compile(`([a-zA-Z&*=\(\)\[\-\_\/\]\>\.\,\<\{\}~^%$#@!\\]+)`)
	return regex.MatchString(stringg)
}																										
																										
func IsEmpty(stringg string) bool {
	if stringg == "" {
		return true
	}
	return false
}
