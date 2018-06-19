package main

import (
	"database/sql"
	"fmt"
	"encoding/csv"
	_ "github.com/lib/pq"
	"os"
	"time"
	"sync"
	
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

func readFile(filename string, comma rune ) [][]string {
	csvFile, err := os.Open(filename) //ioutil.ReadFile("quotes_all.csv")
	check(err)

	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	r.Comma =comma
	r.FieldsPerRecord =-1
	data, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func createTables(db *sql.DB) {
	SQL := `CREATE TABLE IF NOT EXISTS Field ( 
		id 		SERIAL PRIMARY KEY NOT NULL,
		Field1 	TEXT NOT NULL, 
		Field2 TEXT NOT NULL,
		Field3 TEXT NOT NULL);`

	_, err := db.Exec(SQL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created successfully")
}

func insertTable(db *sql.DB, data [][]string) {
	// Insert := fmt.Sprintf(`INSERT INTO Field (Field1, Field2, Field3) VALUES (%s, %s, %s);`, field1, field2, field3)
	fmt.Println(time.Now())

	var wg  sync.WaitGroup
	wg.Add(5)
	LengthData := len(data)
	go1 := LengthData/5
	go2 := go1 + go1
	go3 := go2 +go1
	go4 := go3 +go1
	go5 := LengthData

	go func() {
		defer wg.Done()
		for i := 0; i< go1; i++ {
			_, err := db.Exec("INSERT INTO Field (Field1, Field2, Field3) VALUES ($1, $2, $3);", data[i][0], data[i][1], data[i][2])
			if err != nil {
				
				// panic(err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := go1; i< go2; i++ {
			_, err := db.Exec("INSERT INTO Field (Field1, Field2, Field3) VALUES ($1, $2, $3);", data[i][0], data[i][1], data[i][2])
			if err != nil {
				
				// panic(err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := go2; i< go3; i++ {
			_, err := db.Exec("INSERT INTO Field (Field1, Field2, Field3) VALUES ($1, $2, $3);", data[i][0], data[i][1], data[i][2])
			if err != nil {
				
				// panic(err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := go3; i< go4; i++ {
			_, err := db.Exec("INSERT INTO Field (Field1, Field2, Field3) VALUES ($1, $2, $3);", data[i][0], data[i][1], data[i][2])
			if err != nil {
				
				// panic(err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := go4; i< go5; i++ {
			_, err := db.Exec("INSERT INTO Field (Field1, Field2, Field3) VALUES ($1, $2, $3);", data[i][0], data[i][1], data[i][2])
			if err != nil {
				
				// panic(err)
			}
		}
	}()
	
	wg.Wait() 
	// for i := 0; i< LengthData; i++ {
	// 	_, err := db.Exec("INSERT INTO Field (Field1, Field2, Field3) VALUES ($1, $2, $3);", data[i][0], data[i][1], data[i][2])
	// 	if err != nil {
			
	// 		// panic(err)
	// 	}
	// }
	fmt.Println(time.Now())
}

func main() {
	db := dbConnect()
	data := readFile("quotes_all.csv", ';')
	createTables(db)
	insertTable(db, data)
}
