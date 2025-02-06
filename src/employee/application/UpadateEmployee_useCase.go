package application

import "Store/src/employee/domain"

type UpdateEmployee struct {
	repository domain.IEmployee
}

func NewUpdateEmployee(repository domain.IEmployee) *UpdateEmployee {
	return &UpdateEmployee{repository}
}

func (u *UpdateEmployee) Run(employee domain.Employee) error {
	return u.repository.Update(employee)
}
