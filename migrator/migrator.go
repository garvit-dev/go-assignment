package migrator

import (
	db "Assignment/db"
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	"log"

)

func Createtable() {
	var pg_db = db.Connection()
	
	if err := CreateEmployee(pg_db) ; err != nil {
		log.Printf("Error while creating table employee, Error : %v\n", err)
	}
}
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
