package employee

import (
	"log"
	"encoding/json"
	"fmt"
	db "Assignment/db"
	aqua "github.com/rightjoin/fuel"
	migrator "Assignment/migrator"
)
type EmployeeService struct {
		aqua.Service `prefix:"employee" root:"/"`
		
		addEmployee         aqua.POST 	 `route:"/"`
		deleteEmployee      aqua.DELETE  `route:"/{id}"`
		updateEmployee      aqua.PUT     `route:"/{id}"`
		listEmployee        aqua.POST  	 `route:"/search/"`
}

func init(){
	migrator.Createtable()
}

func (s *EmployeeService) AddEmployee(e aqua.Aide) string {
		
	var (
		err error
		b []byte
	)

	var pg_db = db.Connection() 
	
	if b, err = json.Marshal(e.Post()); err!=nil {
		fmt.Println("error:", err)
	}
	
	employee := &Employee{}
	if err = json.Unmarshal(b, employee); err != nil {
		log.Printf("Error while unmarshalling , Error : %v", err)
		return "Failure"
	}
	
	if VaddEmployee(pg_db) == true {
		if insertErr := employee.PaddEmployee(pg_db) ; insertErr != nil {
			log.Printf("Error while inserting into employee, Error : %v\n", insertErr)
			return "failure"
		}
			return "sucessfully added"
	}
	return "not vallidate"
		
}

func (s *EmployeeService) ListEmployee(e aqua.Aide) string {
	var (
			err error
			b []byte
		)

	var pg_db = db.Connection()
	 if b, err = json.Marshal(e.Post()); err != nil {
	 	fmt.Println("error:", err)
	 } 
	
	employee := &Employee{}
	if err = json.Unmarshal(b, employee) ; err != nil {
		log.Printf("Error while unmarshalling , Error : %v", err)
		return "Failure"
	}
	if VlistEmployee(employee,pg_db) == true {
		if err = PlistEmployee(employee,pg_db) ; err !=nil{
			log.Printf("error while searching an employee",err)
		}

		return "sucessfully search"
	}
	return "not vallidate"
}
func (s *EmployeeService) DeleteEmployee(id string, e aqua.Aide) string {
	var pg_db = db.Connection()

	if VdeleteEmployee(id,pg_db) == true {
		if err := PdeleteEmployee(id,pg_db) ; err != nil {
			log.Printf("error while deleting an employee",err)
		}
		return "sucessfully deleted"
	}
	return "not validate"

}
func (s *EmployeeService) UpdateEmployee(id string, e aqua.Aide) string {
	var pg_db = db.Connection() 
	if VupdateEmployee(id,pg_db) {
		if err := PupdateEmployee(id, pg_db) ; err != nil {
			log.Printf("error while updating employee",err)
		}
		return "sucessfully updated"
	}
	return "not validate"
}
