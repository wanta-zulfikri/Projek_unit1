package user

import "github.com/wanta-zulfikri/Projek_unit1/entities"

type UserInterface interface {
	FindByUsername(username string) (*entities.User, error)
	GetAllByRoleLimit(role string, offset int) ([]*entities.User, error)
	GetAllByRole(role string) ([]*entities.User, error)
	GetAll() ([]*entities.User, error)
	Create(data *entities.User) error
	Update(data *entities.User, userid int) error
	Delete(userid int) error
}
