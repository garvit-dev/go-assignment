package employee

import (
	"encoding/json"
	"fmt"
	"log"

	pg "github.com/go-pg/pg"
	aqua "github.com/rightjoin/fuel"
)

func vaddEmployee(e aqua.Aide) (bool, *Employee) {
	var (
		err error
		b   []byte
	)
	employee := &Employee{}

	if b, err = json.Marshal(e.Post()); err != nil {
		fmt.Println("error:", err)
	}
	if err = json.Unmarshal(b, employee); err != nil {
		log.Printf("Error while unmarshalling , Error : %v", err)
		return false, employee
	}
	return true, employee
}

func vlistEmployee(e aqua.Aide) (bool, *Employee) {
	var (
		err error
		b   []byte
	)
	employee := &Employee{}
	if b, err = json.Marshal(e.Post()); err != nil {
		fmt.Println("error:", err)
	}

	if err = json.Unmarshal(b, employee); err != nil {
		log.Printf("Error while unmarshalling , Error : %v", err)
		return false, employee
	}
	return true, employee
}
func vupdateEmployee(e aqua.Aide) (bool, *Employee) {
	var (
		err error
		b   []byte
	)
	employee := &Employee{}

	if b, err = json.Marshal(e.Post()); err != nil {
		fmt.Println("error:", err)
	}

	if err = json.Unmarshal(b, employee); err != nil {
		log.Printf("Error while unmarshalling , Error : %v", err)
		return false, employee
	}
	return true, employee
}
func vdeleteEmployee(id string, db *pg.DB) bool {
	var err error
	if err = db.Model(&Employee{}).Where("Id IN (?)", id).Select(); err != nil {
		fmt.Printf("Email id not exists", err)
		return false
	}
	return true

}
