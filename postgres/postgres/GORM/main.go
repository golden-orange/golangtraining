package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	// email    = "findme@scottlaing.ca"
	port     = 5432
	user     = "postgres"
	dbname   = "postgres"
	password = "password"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Logging with GORM
	db.LogMode(true)

	// Create a new record with GORM (this is incomplete... see GORM2)
	//	name, email := get the name and email from the user
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }
	if err = db.Create(u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)

	db.AutoMigrate(&User{})
}
