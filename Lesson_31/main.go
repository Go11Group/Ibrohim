package main

import (
	"Mod_Index/model"
	"Mod_Index/postgres"
	"fmt"
	"math/rand"
	"github.com/go-faker/faker/v4"
)


func InsertData(rep *postgres.ProductRepo) error {
	for i := 0; i < 1000000; i++ {
		err := rep.CreateProduct(
			model.Product{
			Name: faker.FirstName(),
			Category: faker.ChineseLastName(),
			Cost: rand.Intn(10000) + 1},
		)
		if err != nil {
			fmt.Println("error inserting new data into the product table")
		}
		if i % 10000 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}


func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	prRepo := postgres.NewProductRepo(db)
	
	/*err = InsertData(prRepo) // inserting a million number of products (BIR marta ishlatilsin)
	if err != nil {
		panic(err)
	}*/

	/*// err = prRepo.CreateIndexCost()
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Index for cost created successfully")
	// }
	q := "select * from product where cost = $1"
	p := []interface{}{5} // 6000, 800, 5
	err = prRepo.ExplainQuery(q, p)
	if err != nil {
		panic(err)
	}*/

	// err = prRepo.CreateIndexNameCategory()
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Index for name and category created successfully")
	// }
	q := "select * from product where category = $1 and name = $2"
	p := []interface{}{"千凡", "Alva"} // [[Lavada,巧兰], [Royal, 夜南], [千凡], Alva]
	err = prRepo.ExplainQuery(q, p)
	if err != nil {
		panic(err)
	}
}