package application

import "Store/src/employee/domain"

type GetAllEmployee struct {
	db domain.IEmployee
}

func NewGetAllEmployee(db domain.IEmployee) *GetAllEmployee {
	return &GetAllEmployee{db: db}
}

func (uc *GetAllEmployee) Run() ([]map[string]interface{}, error) {
	employees, err := uc.db.GetAll()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, employee := range employees {
		result = append(result, employee.ToJSON())
	}
	return result, nil
}
