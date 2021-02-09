package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
	Gender string
	Email string
}

func main()  {
	
}

func Initialize()  {
	db, err = gorm.Open("postgres", "users_db.db")
	if err != nil{
		fmt.Println(err.Error())
		panic("Unable to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{
		Model:  gorm.Model{},
		Name:   "Gill Erick",
		Gender: "Male",
		Email:  "ogayogill95@gmail.com",
	})

}
func AllUsers(){
	db, err = gorm.Open("postgres", "users_db.db")
	if err != nil{
		panic("Could not open the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
}

func AddUser()  {
	
}

func UpdateUser()  {
	
}

func DeleteUser()  {
	
}