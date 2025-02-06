package domain

type IEmployee interface {
	Create(employee Employee) (int32, error)
	GetAll() ([]Employee, error)
	GetByID(id int32) (Employee, error)
	Update(employee Employee) error
	Delete(id int32) error
}
