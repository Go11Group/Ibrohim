package main

import (
	"Exam/tasks"
	"encoding/json"
	"os"
)

type Input struct {
	Slc1 []int
	Text2 string
	Num int
}
type Answer struct {
	Sum int
	Titled string
	Unum int
}

func main() {
	f,err := os.OpenFile("files/input.json",os.O_RDONLY,0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := Input{}
	de := json.NewDecoder(f)
	err = de.Decode(&input)
	if err != nil {
		panic(err)
	}

	res1, err := tasks.SumOfSlice(input.Slc1)
	if err != nil {
		panic(err)
	}
	res2, err := tasks.TextToTitle(input.Text2)
	if err != nil {
		panic(err)
	}
	num := make(chan int)
	go tasks.Abs(num)
	res3 := input.Num
	ans := Answer{
		Sum: res1,
		Titled: res2,
		Unum: res3}
	
	file, err := os.OpenFile("files/output.json",os.O_WRONLY,0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dt,err := json.Marshal(ans)
	if err != nil {
		panic(err)
	}
	file.Write(dt)
	println("The data has been read and written successfully!")
}