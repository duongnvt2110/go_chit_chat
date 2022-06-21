package psql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
) 

const (
	Host = "localhost"
	Port = 5432
	Username = "root"
	Password = "root"
	Database = "postgres"
	Sslmode = "disable"
)

var (
	Psql *sql.DB
)

func init(){
	var err error
	dnsPsql := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",Username,Password,Host,Port,Database,Sslmode)
	// postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	Psql, err = sql.Open("postgres",dnsPsql)
	if err != nil {
		log.Println("Database connect have some problem")
		return
	}
	if err = Psql.Ping(); err != nil {
		log.Println("Database connect have some problem")
		return 
	}
	log.Println("Database successfully configrured")
}
