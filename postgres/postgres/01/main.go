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

	// first version without the // ID
	_, err = db.Exec(`
    INSERT INTO users(name, email)
    VALUES($1, $2)`,
		"Ringo Starkey", "ringo@beatles.com")
	if err != nil {
		panic(err)
	}

	// second version that returns the ID
	var id int
	row := db.QueryRow(`
        INSERT INTO users(name, email)
        VALUES($1, $2) RETURNING id`,
		"Bob Andug", "bob@email.com")

	// querying user by ID
	var name, email string
	row = db.QueryRow(`
        SELECT id, name, email
        FROM users
        WHERE id=$1`, 1)

	err = row.Scan(&id, &name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println("ID:", id, "Name:", name, "Email:", email)

	// Querying users by email address or ID
	rows, err := db.Query(`
      SELECT id, name, email
      FROM users
      WHERE email = $1
      OR ID > $2`,
		"ringo@beatles.com", 3)
	if err != nil {
		panic(err)
	}



	defer db.Close()

	//   sqlStatement := `
	// INSERT INTO users (age, email, first_name, last_name)
	// VALUES ($1, $2, $3, $4)
	// RETURNING id`
	//   id := 0
	//   err = db.QueryRow(sqlStatement, 27, "gh@beatles.com", "George", "Harrison").Scan(&id)
	//
	//   if err != nil {
	//     panic(err)
	//   }
	//
	//   fmt.Println("New record ID is:", id)
	//
}
