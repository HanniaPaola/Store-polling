package application

import "Store/src/employee/domain"

type GetEmployeeByID struct {
	db domain.IEmployee
}

func NewGetEmployeeByID(db domain.IEmployee) *GetEmployeeByID {
	return &GetEmployeeByID{db: db}
}

func (uc *GetEmployeeByID) Run(id int32) (domain.Employee, error) {
    return uc.db.GetByID(id)
}
