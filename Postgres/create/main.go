package create

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "tvducmt"
	dbname   = "csv_db"
)


func dbConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}


func createTables(db *sql.DB) {
	SQL := `CREATE TABLE IF NOT EXISTS Field ( 
		id 		SERIAL PRIMARY KEY NOT NULL,
		Field1 	TEXT NOT NULL, 
		Field2 TEXT NOT NULL UNIQUE,
		Field3 TEXT NOT NULL);`

	_, err := db.Exec(SQL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created successfully")
}

func main() {
	db := dbConnect()
	createTables(db)
	
}
