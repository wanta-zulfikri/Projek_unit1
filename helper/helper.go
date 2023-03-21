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
	if len(stringg) < 1 {
		return true
	}
	return false
}
