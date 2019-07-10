package employee

import (
	migrator "Assignment/migrator"
	"log"

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

func (s *EmployeeService) AddEmployee(e aqua.Aide) string {

	var (
		employee *Employee
		val      bool
		db       *pg.DB
	)
	if db = DB.Connection(); db == nil {
		return "DB Connection failed"
	}

	if val, employee = vaddEmployee(e); val == true {
		if insertErr := employee.paddEmployee(db); insertErr != nil {
			log.Printf("Error while inserting into employee, Error : %v\n", insertErr)
			return "failure"
		}
		return "sucessfully added"
	}
	return "not vallidate"

}

func (s *EmployeeService) ListEmployee(e aqua.Aide) *ListofEmployee {
	var (
		db       *pg.DB
		employee *Employee
		list     *ListofEmployee
		val      bool
		err      error
	)
	// listResponse := ListofEmployee{}

	if db = DB.Connection(); db == nil {
		log.Printf("DB Connection failed")
		// listResponse.Response = "False"
		return list
	}
	if val, employee = vlistEmployee(e); val == true {
		if err, list = plistEmployee(employee, db); err != nil {
			log.Printf("error while searching an employee", err)
		}

		return list
	}
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
	if db = DB.Connection(); db == nil {
		log.Printf("DB Connection failed")
		// listResponse.Response = "False"
		return list
	}

	if val, employee = vupdateEmployee(e); val == true {
		if err, list = pupdateEmployee(employee, db); err != nil {
			log.Printf("error while updating employee", err)
		}
		return list
	}
	return list
}
func (s *EmployeeService) DeleteEmployee(id string, e aqua.Aide) string {
	var db *pg.DB

	if db = DB.Connection(); db == nil {
		return "DB Connection failed"
	}

	if vdeleteEmployee(id, db) == true {
		if err := pdeleteEmployee(id, db); err != nil {
			log.Printf("error while deleting an employee", err)
		}
		return "sucessfully deleted"
	}
	return "email id not exists"

}
