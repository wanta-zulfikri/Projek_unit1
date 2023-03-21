package helper

import "github.com/wanta-zulfikri/Projek_unit1/entities"

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
