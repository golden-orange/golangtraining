package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	email    = "findme@scottlaing.ca"
	port     = 5432
	user     = "postgres"
	dbname   = "postgres"
	password = "password"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Creating orders to query

	var id int

	for i := 1; i < 6; i++ {
		// create some fake data
		userId := 1
		if i > 3 {
			userId = 2
		}

		amount := 1000 * i
		description := fmt.Sprintf("USB-C Adapter x%d", i)

		err = db.QueryRow(`
      INSERT INTO orders (user_id, amount, description)
      VALUES ($1, $2, $3)
      RETURNING id`,
			userId, amount, description).Scan(&id)

		if err != nil {
			panic(err)
		}
		fmt.Println("Created an order with the ID:", id)
	}
	defer db.Close()
}
