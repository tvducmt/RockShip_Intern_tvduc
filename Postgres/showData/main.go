package main

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

type quoteInfo struct {
	id int 
	Field1 string
	Field2 string
	Field3 string
}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}



func showAllTables(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Field")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	objs := make([]quoteInfo, 0)
	for rows.Next() {
		field := quoteInfo{}
		err := rows.Scan(&field.id, &field.Field1, &field.Field2, &field.Field3)
		if err != nil {
			panic(err)
		}
		objs = append(objs, field)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, obj := range objs {
		fmt.Println( obj.Field1, obj.Field1, obj.Field1)
	}
}



func main() {
	db := dbConnect()
	
	 showAllTables(db)
	

}
