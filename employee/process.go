package employee

import (
	"log"

	pg "github.com/go-pg/pg"
)

func (employee *Employee) paddEmployee(db *pg.DB) error {
	if err := db.Insert(employee); err != nil {
		return err
	}
	return nil
}

func plistEmployee(e *Employee, db *pg.DB) (error, *ListofEmployee) {
	var (
		employee Employee
	)
	list := &ListofEmployee{}
	if err := db.Model(&employee).Where("Email = ? ", e.Email).Select(); err != nil {
		return err, list
	}
	list.Name = employee.Name
	list.Id = employee.Id
	list.Email = employee.Email
	list.Password = employee.Password
	list.Phone = employee.Phone

	log.Printf("the employee of particular email id is", employee)
	return nil, list
}

func pdeleteEmployee(id string, db *pg.DB) error {
	if _, err := db.Model(&Employee{}).Where("Id IN (?)", id).Delete(); err != nil {
		return err
	}
	return nil
}

func pupdateEmployee(e *Employee, db *pg.DB) (error, *ListofEmployee) {
	/*var employee Employee*/
	list := &ListofEmployee{}
	employee := &Employee{}
	if _, updateErr := db.Model(employee).Set("Name=?", e.Name).Where("id = ?", e.Id).Update(); updateErr != nil {
		return updateErr, list
	}

	list.Name = employee.Name
	list.Id = employee.Id
	list.Email = employee.Email
	list.Password = employee.Password
	list.Phone = employee.Phone

	return nil, list
}
