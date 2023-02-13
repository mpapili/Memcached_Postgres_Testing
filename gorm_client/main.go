package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Pet is the model that represents the "Pets" table
type Pet struct {
	gorm.Model
	Name string
	Type string
	Age  int
}

func main() {
	// Connect to the database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mikes_db password=password, sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Pet{})
	// if age is not set, set it to 5
	db.Model(&Pet{}).Where("age = ?", 0).Update("age", 5)

	fmt.Println("Table 'Pets' created successfully")

	// Can we create "pets"?
	/*
		fmt.Println("Creating  a new pet")
		myPet := Pet{
			Name: "Marta",
			Type: "Cat",
		}

		db.Create(&myPet)
	*/
	// Some "GET" tests:

	// fetch all pets
	var pets []Pet
	db.Find(&pets)
	for _, pet := range pets {
		fmt.Println(pet)
	}

	pets = []Pet{}
	// let's do a filter now
	db.Where("type = ?", "Dog").Find(&pets)
	for _, pet := range pets {
		fmt.Println(pet)
	}

	// let's delete all pets named "Marta"
	// perform your filter:
	db.Where("name = ?", "Marta").Find(&pets)
	for _, pet := range pets {
		db.Delete(&pet)
	}
}
