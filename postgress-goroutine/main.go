package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	// Connect to db
	connStr := "postgres://@localhost/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	goroutines := 300

	// Channel to syncronize goroutines
	channel := make(chan int, goroutines)

	// Generate goroutines to select records and interate them
	for i := 0; i < goroutines; i++ {
		go func(j int) {
			rows, err := db.Query("select name, age from contacts where age = $1", j)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			var name string
			var age int
			for rows.Next() {
				err := rows.Scan(&name, &age)
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			}
			fmt.Println(name, age)
			err = rows.Err()
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			rows.Close()
			channel <- 0
		}(i)
	}

	// Wait for goroutines to finish
	for i := 0; i < goroutines; i++ {
		<-channel
	}
}
