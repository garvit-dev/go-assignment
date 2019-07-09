package main 

import(
	employee "Assignment/employee"
	aqua "github.com/rightjoin/fuel"
)

func main(){

	server := aqua.NewServer()
    server.AddService(&employee.EmployeeService{})
    server.Run()
}