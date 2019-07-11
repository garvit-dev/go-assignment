package employee

import (
	"log"

	migrator "Assignment/migrator"

	DB "Assignment/db"

	pg "github.com/go-pg/pg"
	aqua "github.com/rightjoin/fuel"
)

type EmployeeService struct {
	aqua.Service `prefix:"employee" root:"/"`

	addEmployee    aqua.POST   `route:"/"`
	deleteEmployee aqua.DELETE `route:"/{id}"`
	updateEmployee aqua.PUT    `route:"/update/"`
	listEmployee   aqua.POST   `route:"/search/"`
}

func init() {
	migrator.Createtable()
}

func (s *EmployeeService) AddEmployee(e aqua.Aide) Resp {

	var (
		employee  *Employee
		val       bool
		db        *pg.DB
		insertErr error
	)

	if val, employee = vaddEmployee(e); val == true {
		if db = DB.Connection(); db == nil {
			log.Printf("DB Connection failed")
		}

		if insertErr = employee.paddEmployee(db); insertErr != nil {
			log.Printf("Error while inserting into employee, Error : %v\n", insertErr)
			return buildResponse(insertErr, nil)
		}
		return buildResponse(insertErr, nil)
	}
	defer DB.Close(db)
	return buildResponse(insertErr, nil)

}

func (s *EmployeeService) ListEmployee(e aqua.Aide) *ListofEmployee {
	var (
		db       *pg.DB
		employee *Employee
		list     *ListofEmployee
		val      bool
		err      error
	)

	if val, employee = vlistEmployee(e); val == true {

		if db = DB.Connection(); db == nil {
			log.Printf("DB Connection failed")
			return list
		}

		if err, list = plistEmployee(employee, db); err != nil {
			log.Printf("error while searching an employee", err)
		}

		return list
	}
	defer DB.Close(db)
	return list
}

func (s *EmployeeService) UpdateEmployee(e aqua.Aide) *ListofEmployee {
	var (
		db       *pg.DB
		employee *Employee
		val      bool
		list     *ListofEmployee
		err      error
	)

	if val, employee = vupdateEmployee(e); val == true {
		if db = DB.Connection(); db == nil {
			log.Printf("DB Connection failed")
			return list
		}

		if err, list = pupdateEmployee(employee, db); err != nil {
			log.Printf("error while updating employee", err)
		}
		return list
	}
	defer DB.Close(db)
	return list
}
func (s *EmployeeService) DeleteEmployee(id string, e aqua.Aide) string {
	var db *pg.DB

	if vdeleteEmployee(id, db) == true {
		if db = DB.Connection(); db == nil {
			return "DB Connection failed"
		}

		if err := pdeleteEmployee(id, db); err != nil {
			log.Printf("error while deleting an employee", err)
		}
		return "sucessfully deleted"
	}
	defer DB.Close(db)
	return "email id not exists"

}
