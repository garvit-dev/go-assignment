package migrator

import (

)
type Employee struct{
	Id string  `sql:"id,pk type:serial"`
	Name string `sql:"name"`
	Email  string `sql:"email"`
	Phone_number string `sql:"phonenumber"`
	Password string `sql:"pass"`
}