package db

import(
pg "github.com/go-pg/pg"
)

func Connection() *pg.DB {
		opts:= &pg.Options{
		User : "postgres",
		Password : "12345",
		Addr: "127.0.0.1:5432",
    }
	pg_db := pg.Connect(opts)
	
	return pg_db	
	 
}
