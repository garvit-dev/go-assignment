package employee

import(
pg "github.com/go-pg/pg"
)
type Employee struct{
	Id string  
	Name string
	Email  string
	Phone_number string
	Password string
}

func Connection() *pg.DB{
		opts:= &pg.Options{
		User : "postgres",
		Password : "12345",
		Addr: "127.0.0.1:5432",
    }
	pg_db := pg.Connect(opts)
	return pg_db 
}


