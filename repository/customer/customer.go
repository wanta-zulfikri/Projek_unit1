package customer

import "github.com/wanta-zulfikri/Projek_unit1/entities"

type CustomerInterface interface {
	GetAll() ([]*entities.Customer, error)
	GetWithLimit(offset int) ([]*entities.Customer, error)
	FindByPhone(phonenum string) (*entities.Customer, error)
	Delete(userid int) error
	Update(data *entities.Customer) error
	Create(data *entities.Customer) error
}
