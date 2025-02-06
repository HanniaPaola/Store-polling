package application

import "Store/src/employee/domain"

type DeleteEmployee struct {
	repository domain.IEmployee
}

func NewDeleteEmployee(repository domain.IEmployee) *DeleteEmployee {
	return &DeleteEmployee{repository}
}

func (d *DeleteEmployee) Run(id int32) error {
	return d.repository.Delete(id)
}
