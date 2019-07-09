package employee
	
import (
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	"log"
	)
func CreateEmployee(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	if createErr := db.CreateTable(&Employee{}, opts) ; createErr != nil {
		return createErr
	}
	log.Printf("employee table created\n")
	return nil
}

func (employee *Employee)AddEmployee(db *pg.DB) error {
    if err:= db.Insert(employee) ; err != nil {
    	return err
    }
    return nil 
}

func ListEmployee(s *Employee, db *pg.DB) error {
  	var employee Employee
	if err := db.Model(&employee).Where("Email = ? ", s.Email).Select() ; err != nil {
		return err 
	}
	log.Printf("the employee of particular email id is", employee)
	return nil
}

func DeleteEmployee(id string, db *pg.DB) error {
	if _,err := db.Model(&Employee{}).Where("Id IN (?)", id).Delete() ; err != nil{
		return err
	}
	return nil
}

func UpdateEmployee(id string , db *pg.DB) error{
	/*var employee Employee*/
	if _,updateErr := db.Model(&Employee{}).Set("Name=?","Nidhi").Where("id = ?",id).Update(); updateErr !=nil {
		return updateErr
	}	
	return nil
}
 