package main

import (
	// "User_Gorm/model"
	"User_Gorm/orm/postgres"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/user?sslmode=disable"))
	if err != nil {
		panic(err)
	}

	u := user.UserRepo{DB: db}
	/*err = u.CreateUserTable()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("User table created successfully")
	}*/

	users, err := u.GetAllUsers()
	if err != nil {
		panic(err)
	}
	for _, v := range users {
		fmt.Println(v)
	}

	/*user, err := u.GetUser("John")
	if err != nil {
		panic(err)
	}
	fmt.Println(*user)*/

	/*err = u.Create(model.User {
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
		Password:   "randompassword123",
		Age:        30,
		Field:      "Engineering",
		Gender:     "Male",
		IsEmployee: true})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Data inserted successfully")
	}*/

	/*err = u.Update("John", model.User{
		LastName:   "Smith",
    	Email:      "jane.smith@example.com",
    	Password:   "securepassword456",
    	Age:        25,
    	Field:      "Medicine",
    	Gender:     "Female",
    	IsEmployee: false})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Data updated successfully")
	}*/

	/*err = u.Delete("John")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Data deleted successfully")
	}*/
}