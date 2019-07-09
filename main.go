package main 

import(
	employee "Assignment/employee"s
	aqua "github.com/rightjoin/fuel"
)

func main(){

	server := aqua.NewServer()
    server.AddService(&employee.EmployeeService{})
    server.Run()
}