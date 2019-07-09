package employee
	
import (
	pg "github.com/go-pg/pg"
	"log"
	)
func (employee *Employee)PaddEmployee(db *pg.DB) error {
    if err:= db.Insert(employee) ; err != nil {
    	return err
    }
    return nil 
}

func PlistEmployee(e *Employee, db *pg.DB) error {
  	var employee Employee
	if err := db.Model(&employee).Where("Email = ? ", e.Email).Select() ; err != nil {
		return err 
	}
	log.Printf("the employee of particular email id is", employee)
	return nil
}

func PdeleteEmployee(id string, db *pg.DB) error {
	if _,err := db.Model(&Employee{}).Where("Id IN (?)", id).Delete() ; err != nil{
		return err
	}
	return nil
}

func PupdateEmployee(e *Employee,db *pg.DB) error{
	/*var employee Employee*/

	if _,updateErr := db.Model(&Employee{}).Set("Name=?",e.Name).Where("id = ?",e.Id).Update(); updateErr !=nil {
		return updateErr
	}	
	return nil
}
 