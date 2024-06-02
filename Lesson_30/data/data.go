package data

import "User_Product/models"

func InitialiseData() ([]model.User, []model.Product) {
	productInfo := []model.Product {
		{
		Name: "Wireless Headphones",
		Description: "Experience crystal-clear sound with these sleek wireless headphones.",
		Price: 79.99,
		Stock_quantity: 50,
		},
		
		{
		Name: "Smartphone Case",
		Description: "Protect your smartphone with this stylish and durable case.",
		Price: 19.99,
		Stock_quantity: 100,
		},
		
		{
		Name: "Bluetooth Speaker",
		Description: "Take your music anywhere with this portable Bluetooth speaker.",
		Price: 49.99,
		Stock_quantity: 30,
		},
		
		{
		Name: "Fitness Tracker",
		Description: "Stay active and track your fitness goals with this sleek tracker.",
		Price: 89.99,
		Stock_quantity: 20,
		},
		
		{
		Name: "Coffee Maker",
		Description: "Brew delicious coffee at home with this easy-to-use coffee maker.",
		Price: 39.99,
		Stock_quantity: 15,
		},
		
		{
		Name: "Gaming Mouse",
		Description: "Enhance your gaming experience with this high-performance gaming mouse.",
		Price: 59.99,
		Stock_quantity: 25,
		},
		
		{
		Name: "Portable Charger",
		Description: "Never run out of battery again with this convenient portable charger.",
		Price: 29.99,
		Stock_quantity: 40,
		},
		
		{
		Name: "Wireless Keyboard",
		Description: "Enjoy the freedom of wireless typing with this ergonomic keyboard.",
		Price: 69.99,
		Stock_quantity: 35,
		},
		
		{
		Name: "Smart Watch",
		Description: "Stay connected and organized with this feature-packed smartwatch.",
		Price: 129.99,
		Stock_quantity: 10,
		},
		
		{
		Name: "Digital Camera",
		Description: "Capture your memories in stunning detail with this high-resolution digital camera.",
		Price: 199.99,
		Stock_quantity: 5,
		}}
	
	userInfo := []model.User {
		{Username: "JohnDoe", Email: "johndoe@example.com", Password: "JD@1234"},
		{Username: "JaneSmith", Email: "janesmith@example.com", Password: "JS@5678"},
		{Username: "MikeJohnson", Email: "mikejohnson@example.com", Password: "MJ@abcd"},
		{Username: "EmilyBrown", Email: "emilybrown@example.com", Password: "EB@efgh"},
		{Username: "ChrisWilson", Email: "chriswilson@example.com", Password: "CW@ijkl"},
		{Username: "SarahTaylor", Email: "sarahtaylor@example.com", Password: "ST@mnop"},
		{Username: "AlexMiller", Email: "alexmiller@example.com", Password: "AM@qrst"},
		{Username: "LauraDavis", Email: "lauradavis@example.com", Password: "LD@uvwx"},
		{Username: "KevinMartinez", Email: "kevinmartinez@example.com", Password: "KM@yz12"},
		{Username: "JessicaClark", Email: "jessicaclark@example.com", Password: "JC@3456"}}
	return userInfo, productInfo
}