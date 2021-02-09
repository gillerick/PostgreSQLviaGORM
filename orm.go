package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



var db *gorm.DB
var err error


type User struct {
	gorm.Model
	Name string
	Gender string
	Email string
}


func Initialize()  {
	ConnectDatabase()

	db.AutoMigrate(&User{
		Model:  gorm.Model{},
		Name:   "Gill Erick",
		Gender: "Male",
		Email:  "ogayogill95@gmail.com",
	})

}

//Creates a new user table,drops existing one
func CreateTable(){
	defer db.Close()
	database := ConnectDatabase()
	database.DropTableIfExists(&User{})
	database.CreateTable(&User{})
	fmt.Printf("Table successfully created\n")
}

//Connect database function, returns: db gorm.DB
func ConnectDatabase() *gorm.DB {
	db, err = gorm.Open("postgres", "user=postgres port=5433 password=Thanks001 database=postgres sslmode=disable")
	if err != nil{
		fmt.Println(err.Error())
		panic("Unable to connect to database\n")
	}


	database := db.DB()
	err = database.Ping()

	if err != nil{
		panic(err.Error())
	}
	fmt.Printf("Successfully connected to database\n")
	return db
}


func CreateDB() {
	ConnectDatabase()
	db = db.Exec("CREATE DATABASE users;")
	if db.Error != nil {
		fmt.Println("Unable to create the database")
	}
	fmt.Println("Database successfully created")
}

func main()  {
	Initialize()
	CreateDB()
	CreateTable()

}