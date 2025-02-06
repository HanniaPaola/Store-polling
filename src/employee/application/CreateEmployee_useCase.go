package application

import "Store/src/employee/domain"

type CreateEmployee struct {
	repo domain.IEmployee
}

func NewCreateEmployee(repo domain.IEmployee) *CreateEmployee {
	return &CreateEmployee{repo: repo}
}

func (uc *CreateEmployee) Run(employee *domain.Employee) error {
	id, err := uc.repo.Create(*employee)
	if err != nil {
		return err
	}
	employee.SetID(id)
	return nil
}
