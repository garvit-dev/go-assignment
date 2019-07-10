package migrator

type Employee struct {
	Id       string `sql:"id,pk type:serial"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Phone    string `sql:"phone"`
	Password string `sql:"password"`
}
