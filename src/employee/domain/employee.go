package domain

type Employee struct {
	id       int32
	name     string
	position string
	salary   int // Cambiado a int32
}

// Constructor para crear un nuevo Employee
func NewEmployee(name string, position string, salary int32) *Employee { // Cambiado a int32
	return &Employee{id: 1, name: name, position: position, salary: int(salary)}
}

// Método para devolver un empleado serializable
func (e *Employee) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":       e.id,
		"name":     e.name,
		"position": e.position,
		"salary":   e.salary,
	}
}

// Métodos Get y Set
func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetPosition() string {
	return e.position
}

func (e *Employee) SetPosition(position string) {
	e.position = position
}

func (e *Employee) GetSalary() int { // Cambiado a int32
	return e.salary
}

func (e *Employee) SetSalary(salary int) { // Cambiado a int32
	e.salary = salary
}

func (e *Employee) GetID() int32 {
	return e.id
}

func (e *Employee) SetID(id int32) {
	e.id = id
}
