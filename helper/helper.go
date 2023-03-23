package helper

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/entities"
	"github.com/wanta-zulfikri/Projek_unit1/repository/transaksi"
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
func PrintData(datas ...interface{}) {
	if datas, ok := datas[0].([]*entities.User); ok {
		fmt.Println("Berikut List Pegawai: ")
		for i, val := range datas {
			fmt.Printf("%d. Nama Pegawai : %s , Password Pegawai: %s\n", i+1, val.Username, val.Password)
		}
	}
	if datas, ok := datas[0].([]*entities.Customer); ok {
		fmt.Println("Berikut List Customer: ")
		for i, val := range datas {
			fmt.Printf("%d. Nama Customer : %s\nDetail :\nAlamat: %s\nNo HP : %s\n\n", i+1, val.Nama, val.Alamat, val.NoHp)
		}
	}
	if datas, ok := datas[0].([]*entities.Produk); ok {
		fmt.Println("Berikut List Produk: ")
		for i, val := range datas {
			if !val.Nama_Pengganti.Valid {
				val.Nama_Pengganti.String = "Belum Ada"
			}
			fmt.Printf("%d. Nama Produk : %s\nQty : %d\nHarga: %d\nDibuat oleh: %s\nTerakhir Update:%s\n\n", i+1, val.Nama_produk, val.Qty, val.Price, val.Nama_pembuat, val.Nama_Pengganti.String)
		}
	}

	if datas1, ok := datas[0].([]*entities.Transaksi); ok {
		if interfacet, ok := datas[1].(transaksi.TransaksiInterface); ok {
			fmt.Println("Berikut List Transaksi")
			for i, val := range datas1 {
				fmt.Printf("%d. Id Transaksi: %d\nTanggal Transaksi: %s\nNama Pelanggan:%s\nDetail Barang:\n", i+1, val.Id, val.Tanggal.String(), val.CustomerName)
				items, _ := interfacet.GetListItemByid(val.Id)
				for j, val2 := range items {
					fmt.Printf("\t%d. Nama Produk: %s\n\tJumlah Produk: %d\n\tHarga Produk:%d\n", j+1, val2.ProdukName, val2.Qty, val2.Price)
				}
				fmt.Print("\n\n")
			}
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
	if len(stringg) == 0 {
		return true
	}
	regex, _ := regexp.Compile(`([a-zA-Z&*=\"\'\^\(\)\[\-\_\/\]\>\.\,\<\{\}~^%$#@!\\]+)`)
	return regex.MatchString(stringg)
}

func IsEmpty(stringg string) bool {
	if stringg == "" {
		return true
	}
	return false
}
func IfAlreadyTaken(data string, input string) bool {
	regex, _ := regexp.Compile(fmt.Sprintf(`([%s]+)`, data))
	return regex.MatchString(input)
}
