package test

import (
	"log"
	inventory "module/assignment/inventory"
	"module/assignment/products"
	"testing"
	//d "module/assignment/database"
)

func TestAddProduct(t *testing.T) {
	inv := &inventory.Inventory{
		Products: make(map[int]products.Product), // Initialize the Products map
	}
	//	inv := inventory.NewInventory()

	p1 := products.Product{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"}

	expected := 1

	inv.AddProduct(p1)

	if expected != len(inv.Products) {
		t.Error("Expected :", expected, " Result :", len(inv.Products))
	}

}

func TestRemoveProduct(t *testing.T) {
	inv := &inventory.Inventory{
		Products: make(map[int]products.Product), // Initialize the Products map
	}
	p1 := products.Product{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"}
	p2 := products.Product{ID: 2, Name: "Laptops", Price: 999.00, Quantity: 15, Category: "Software"}

	inv.AddProduct(p1)
	inv.AddProduct(p2)

	//expected := false

	err := inv.RemoveProduct(2)
	if err != nil {
		log.Fatal("RemoveProduct function Failed ", err)
	}

	_, exists := inv.Products[2]

	if exists {
		t.Error("Expected Deletion But found the Element ", inv.Products[2])
	}

}

func TestFindProductByName(t *testing.T) {
	inv := &inventory.Inventory{
		Products: make(map[int]products.Product), // Initialize the Products map
	}
	p1 := products.Product{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"}
	p2 := products.Product{ID: 2, Name: "Laptops", Price: 999.00, Quantity: 15, Category: "Software"}

	inv.AddProduct(p1)
	inv.AddProduct(p2)

	expected := "Laptop"

	res, err := inv.FindProductByName("Laptop")

	if err != nil {
		log.Fatalf("Find product by name failed %s", err)
	}

	if expected != res.Name {
		t.Error("Expected :", expected, " Result :", res.Name)
	}

}

func TestTotalValue(t *testing.T) {
	inv := &inventory.Inventory{
		Products: make(map[int]products.Product), // Initialize the Products map
	}
	//inv := inventory.Inventory

	p1 := products.Product{ID: 1, Name: "Laptop", Price: 10.0, Quantity: 10, Category: "Electronics"}
	p2 := products.Product{ID: 2, Name: "Laptops", Price: 1.0, Quantity: 15, Category: "Software"}

	inv.AddProduct(p1)
	inv.AddProduct(p2)

	expected := float64((float64(p1.Quantity) * (p1.Price)) + (float64(p2.Quantity) * (p2.Price)))

	res := inv.TotalValue()

	if expected != res {
		t.Error("Expected :", expected, " Result :", res, "Count is :", len(inv.Products))
	}

}
