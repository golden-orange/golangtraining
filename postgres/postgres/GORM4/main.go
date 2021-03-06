// Querying relational data with GORM 8 4 9
// Preloading user orders
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
	Name   string
	Email  string  `gorm:"not null;unique_index"`
	Orders []Order // return user with user's associated orders
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
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
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var user User
	db.Preload("Orders").First(&user)
	if db.Error != nil {
		panic(db.Error)
	}

	fmt.Println("Email:", user.Email)
	fmt.Println("Number of orders:", len(user.Orders))
	fmt.Println("Orders:", user.Orders)
}
