package db

import (
	"fmt"

	pg "github.com/go-pg/pg"
)

func Connection() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "12345",
		Addr:     "127.0.0.1:5432",
	}
	db := pg.Connect(opts)

	return db

}

//Close database connection.
func Close(db *pg.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Database Connection Closed Success")
}
