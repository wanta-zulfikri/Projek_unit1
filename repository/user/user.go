package user

import "github.com/wanta-zulfikri/Projek_unit1/entities"

type UserInterface interface {
	FindByUsername(username string) (entities.User, error)
}
