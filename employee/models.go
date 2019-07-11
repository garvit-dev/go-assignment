package employee

// Request model
type Employee struct {
	Id       string `sql:"id,pk type:serial"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Phone    string `sql:"phone"`
	Password string `sql:"password"`
}

//List of Employee based on email
type ListofEmployee struct {
	Id       string `sql:"id,pk type:serial"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Phone    string `sql:"phone"`
	Password string `sql:"password"`
}

//Resp : Response Model
type Resp struct {
	Data   interface{}
	Msg    string
	Error  error
	Status int
}
