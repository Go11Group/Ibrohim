package main

import (
	// "User_Product/data"
	model "User_Product/models"
	"User_Product/postgres"
	"fmt"
	// "math/rand"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	uRepo := postgres.NewUserRepo(db)
	// pRepo := postgres.NewProductRepo(db)
	upRepo := postgres.NewUserProductRepo(db)

	/*userInfo, productInfo := data.InitialiseData()
	for _, u := range userInfo {
		if err := uRepo.CreateUser(u); err != nil { // inserts data into users table
			fmt.Println("Error creating user:", err)
			continue
		}
	}
	for _, p := range productInfo {
		if err := pRepo.CreateProduct(p); err != nil { // inserts data into products table
			fmt.Println("Error creating product:", err)
			continue
		}
	}*/
	
	users, err := uRepo.GetUser(model.User{}) // retrieves data from users table
	if err != nil {
		fmt.Println("Error retrieving users:", err)
	}
	/*products, err := pRepo.GetProduct(model.Product{}) // retrieves data from products table
	if err != nil {
		fmt.Println("Error retrieving products:", err)
	}

	for _,u := range users {
		numProducts := rand.Intn(len(products)) + 1
		rand.Shuffle(len(products), func(i, j int) {
			products[i], products[j] = products[j], products[i]
		})

		for i := 0; i < numProducts; i++ {
			qu := rand.Intn(5) + 1
			randProduct := products[rand.Intn(len(products))]
			err := upRepo.AddProductToUser(u.ID, randProduct.ID, qu)
			if err != nil {
				fmt.Println("Error adding product to user:", err)
				continue
			}
		}
	}*/

	printUsers(upRepo, users)
}

func printUsers(upRepo *postgres.UserProductRepo, users []model.User) {
	for _,u := range users {
		products, err := upRepo.GetUserProducts(u.ID)
		if err != nil {
			fmt.Printf("Error retrieving products for user %d: %v\n", u.ID, err)
        	continue
		}
		fmt.Printf("Products for user %d:\n", u.ID)
		for _,p := range products {
			fmt.Printf("- %s: %s (Stock Quantity: %d)\n", p.Name, p.Description, p.Stock_quantity)
		}
	}
}