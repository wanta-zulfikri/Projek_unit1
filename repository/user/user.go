package user

import "github.com/wanta-zulfikri/Projek_unit1/entities"

type UserInterface interface {
	FindByUsername(username string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Create(data *entities.User) error
	Update(data *entities.User, userid int) error
	Delete(userid int) error
}
