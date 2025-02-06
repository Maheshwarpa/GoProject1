package Database

import (
	//"fmt"
	"math/rand"
	inventory "module/assignment/Inventory"
	Products "module/assignment/products"
	//"time"
)

type ProductStorage interface {
	Save(p Products.Product) error
	GetByID(id int) (*Products.Product, error)
	Delete(id int) error
}

func LoadData(p *inventory.Inventory) {
	products := []Products.Product{
		{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"},
		{ID: 2, Name: "Smartphone", Price: 699.50, Quantity: 25, Category: "Electronics"},
		{ID: 3, Name: "Desk-Chair", Price: 149.99, Quantity: 50, Category: "Furniture"},
		{ID: 4, Name: "Coffee-Maker", Price: 89.99, Quantity: 30, Category: "Appliances"},
		{ID: 5, Name: "Wireless-Mouse", Price: 29.99, Quantity: 100, Category: "Accessories"},
		{ID: 6, Name: "Headphones", Price: 59.99, Quantity: 40, Category: "Electronics"},
		{ID: 7, Name: "External-Hard-Drive", Price: 120.99, Quantity: 20, Category: "Electronics"},
		{ID: 8, Name: "Backpack", Price: 49.99, Quantity: 35, Category: "Accessories"},
		{ID: 9, Name: "Electric-Kettle", Price: 39.99, Quantity: 60, Category: "Appliances"},
		{ID: 10, Name: "Running-Shoes", Price: 89.99, Quantity: 75, Category: "Footwear"},
		{ID: 11, Name: "Smartwatch", Price: 199.99, Quantity: 15, Category: "Electronics"},
		{ID: 12, Name: "Bluetooth-Speaker", Price: 79.99, Quantity: 45, Category: "Electronics"},
		{ID: 13, Name: "Gaming-Keyboard", Price: 99.99, Quantity: 25, Category: "Gaming"},
		{ID: 14, Name: "4K-Monitor", Price: 329.99, Quantity: 10, Category: "Electronics"},
		{ID: 15, Name: "Power-Bank", Price: 45.99, Quantity: 55, Category: "Electronics"},
		{ID: 16, Name: "Smart-TV", Price: 499.99, Quantity: 20, Category: "Electronics"},
		{ID: 17, Name: "Office-Desk", Price: 199.99, Quantity: 10, Category: "Furniture"},
		{ID: 18, Name: "LED-Table-Lamp", Price: 29.99, Quantity: 50, Category: "Home Decor"},
		{ID: 19, Name: "Electric-Toothbrush", Price: 69.99, Quantity: 30, Category: "Personal Care"},
		{ID: 20, Name: "Digital-Camera", Price: 599.99, Quantity: 5, Category: "Electronics"},
		{ID: 21, Name: "Portable-Air-Conditioner", Price: 349.99, Quantity: 8, Category: "Appliances"},
		{ID: 22, Name: "Mechanical-Watch", Price: 150.99, Quantity: 12, Category: "Accessories"},
		{ID: 23, Name: "Sunglasses", Price: 79.99, Quantity: 50, Category: "Fashion"},
		{ID: 24, Name: "Treadmill", Price: 899.99, Quantity: 3, Category: "Fitness"},
		{ID: 25, Name: "Wireless-Earbuds", Price: 129.99, Quantity: 40, Category: "Electronics"},
		{ID: 26, Name: "Yoga-Mat", Price: 25.99, Quantity: 100, Category: "Fitness"},
		{ID: 27, Name: "Espresso-Machine", Price: 249.99, Quantity: 15, Category: "Appliances"},
		{ID: 28, Name: "Electric-Scooter", Price: 699.99, Quantity: 7, Category: "Transport"},
		{ID: 29, Name: "Fitness-Tracker", Price: 99.99, Quantity: 30, Category: "Fitness"},
		{ID: 30, Name: "Leather-Wallet", Price: 39.99, Quantity: 60, Category: "Accessories"},
	}

	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		randomNumber := rand.Intn(30) // Generates a number between 0 and 30
		p.AddProduct(products[randomNumber])
		//fmt.Println(randomNumber)
	}
}
