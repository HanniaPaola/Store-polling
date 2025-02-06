package infraestructureEmployee

import (
	"fmt"
	"Store/src/employee/domain"

	"github.com/go-mysql-org/go-mysql/client"
)

type MySQLEmployee struct {
	Conn *client.Conn
}

func (mysql *MySQLEmployee) Create(employee domain.Employee) (int32, error) {
	query := "INSERT INTO employees (name, position, salary) VALUES (?, ?)"
	stmt, err := mysql.Conn.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(employee.GetName(), employee.GetSalary())
	if err != nil {
		return 0, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	lastInsertId := int32(result.InsertId)
	return lastInsertId, nil
}

func (mysql *MySQLEmployee) GetAll() ([]domain.Employee, error) {
	query := "SELECT id, name, position, salary FROM employees"
	rows, err := mysql.Conn.Execute(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo empleados: %v", err)
	}

	var employees []domain.Employee

	fmt.Printf("Número de filas obtenidas: %d\n", len(rows.Values))

	for _, row := range rows.Values {
		id := row[0].AsInt64()
		name := string(row[1].AsString())
		position := string(row[2].AsString())
		salary := float32(row[3].AsFloat64())

		fmt.Printf("Empleado: ID=%d, Name=%s, Position=%s, Salary=%.2f\n", id, name, position, salary)

		employee := domain.NewEmployee(name,position, int32(salary))
		employee.SetID(int32(id))
		employees = append(employees, *employee)
	}

	if len(employees) == 0 {
		fmt.Println("No se encontraron empleados")
	}

	return employees, nil
}

func (mysql *MySQLEmployee) GetByID(id int32) (domain.Employee, error) {
	query := "SELECT id, name, salary FROM employees WHERE id = ?"
	stmt, err := mysql.Conn.Prepare(query)
	if err != nil {
		return domain.Employee{}, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(id)
	if err != nil {
		return domain.Employee{}, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	fmt.Printf("Filas obtenidas para ID=%d: %d\n", id, len(result.Values))

	if len(result.Values) == 0 {
		return domain.Employee{}, fmt.Errorf("empleado con ID %d no encontrado", id)
	}

	row := result.Values[0]
	idFromDB := row[0].AsInt64()
	name := string(row[1].AsString())
	position := string(row[2].AsString())
	salary := float32(row[3].AsFloat64())

	fmt.Printf("Empleado encontrado: ID=%d, Name=%s, Position=%s, Salary=%.2f\n", idFromDB, name, position, salary)

	// Crea una instancia del empleado y lo devuelve.
	employee := domain.NewEmployee(name, position, int32(salary)) // Asegúrate de tener 'position'
	employee.SetID(int32(idFromDB))
	return *employee, nil
}

func (mysql *MySQLEmployee) Update(employee domain.Employee) error {
	query := "UPDATE employees SET name = ?, salary = ? WHERE id = ?"
	stmt, err := mysql.Conn.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(employee.GetName(), employee.GetSalary(), employee.GetID())
	if err != nil {
		return fmt.Errorf("error actualizando empleado: %v", err)
	}

	if result.AffectedRows == 0 {
		return fmt.Errorf("empleado con ID %d no encontrado", employee.GetID())
	}

	return nil
}

func (mysql *MySQLEmployee) Delete(id int32) error {
	query := "DELETE FROM employees WHERE id = ?"
	stmt, err := mysql.Conn.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(id)
	if err != nil {
		return fmt.Errorf("error eliminando empleado: %v", err)
	}

	if result.AffectedRows == 0 {
		return fmt.Errorf("empleado con ID %d no encontrado", id)
	}

	return nil
}
