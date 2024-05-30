package main

import (
	model "Stu_Cou/models"
	"Stu_Cou/storage/postgres"
	"fmt"
	"time"
	"github.com/google/uuid"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database successfully")
	defer db.Close()

	crepo := postgres.NewCourseRepo(db)
	srepo := postgres.NewStudentRepo(db)
	
	allCourses, err := crepo.GetAllCourses()
	if err != nil {
		panic(err)
	}
	fmt.Println("All courses:")
	for _, course := range allCourses {
		fmt.Printf("ID: %s, Name: %s\n", course.ID, course.Name)
	}

	allStudents, err := srepo.GetAllStudents()
	if err != nil {
		panic(err)
	}
	fmt.Println("All students:")
	for _, student := range allStudents {
		fmt.Printf("ID: %s, Name: %s, Age: %d\n", student.ID, student.Name, student.Age)
	}

	newStudent := model.Student {
		ID: uuid.NewString(),
		Name: "Chris Evans",
		Age: 20,
	}
	err = srepo.Create(newStudent)
	if err != nil {
		panic(err)
	}
	fmt.Println("New student inserted successfully!")
	time.Sleep(time.Second * 15)

	err = srepo.Update(newStudent.ID, "Tom Holland", 25)
	if err != nil {
		panic(err)
	}
	fmt.Println("Student updated successfully")
	time.Sleep(time.Second * 15)

	err = srepo.Delete(newStudent.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Student deleted successfully")
}