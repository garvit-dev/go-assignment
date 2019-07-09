package employee
	
import (
	pg "github.com/go-pg/pg"
	// "log"
	)
type Employee struct{
	Id string  
	Name string
	Email  string
	Phone_number string
	Password string
}
func VaddEmployee(db *pg.DB) bool {
   return true 
}

func VlistEmployee(s *Employee, db *pg.DB) bool {
	return true
}

func VdeleteEmployee(id string, db *pg.DB) bool {
	return true 
}

func VupdateEmployee( db *pg.DB) bool{
	return true
}
 