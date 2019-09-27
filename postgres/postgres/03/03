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

	// Querying users by email address or ID
	var id int
	var name, email string

	rows, err := db.Query(`
      SELECT id, name, email
      FROM users
      WHERE email = $1
      OR ID > $2`,
		"ringo@beatles.com", 7)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created an order with the ID:", id)

	for rows.Next() {
		rows.Scan(&id, &name, &email)
		fmt.Println("ID:", id, "Name:", name, "Email:", email)
	}

	defer db.Close()
}
