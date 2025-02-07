package test

import (
	"fmt"
	"log"
	"module/assignment/products"
	"testing"
)

/*

	CustomerList = []Customer{
		{Custid: 1, Fname: "John", Lname: "Doe", Email: "john.doe@example.com", Age: 28, Pnum: "1234567890", Bal: 5000},
		{Custid: 2, Fname: "Jane", Lname: "Smith", Email: "jane.smith@example.com", Age: 34, Pnum: "0987654321", Bal: 10000},
		{Custid: 3, Fname: "Alice", Lname: "Johnson", Email: "alice.johnson@example.com", Age: 42, Pnum: "1122334455", Bal: 7500},
		{Custid: 4, Fname: "Bob", Lname: "Brown", Email: "bob.brown@example.com", Age: 25, Pnum: "6677889900", Bal: 3000},
		{Custid: 5, Fname: "Charlie", Lname: "Davis", Email: "charlie.davis@example.com", Age: 30, Pnum: "5566778899", Bal: 9000},
	}

	expected := false
	var result bool
	result = CheckAcc(6)

	if expected != result {
		t.Error(" Expected :", expected, " REsult: ", result)
	}
*/

func TestUpdatePrice(t *testing.T) {
	p1 := products.Product{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"}

	expected := 1000.0

	p1.UpdatePrice(expected)

	if expected != p1.Price {
		t.Error(" Expected :", expected, " Result: ", p1.Price)
	}
}

func TestSell(t *testing.T) {
	p1 := products.Product{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"}

	expected := 5

	err := p1.Sell(5)

	if err != nil {
		log.Fatal("Sell Test Case Failed")
	}

	if expected != p1.Quantity {
		t.Error(" Expected :", expected, " Result: ", p1.Quantity)
	}

}

func TestRestock(t *testing.T) {
	p1 := products.Product{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"}

	expected := 15

	p1.Restock(5)

	if expected != p1.Quantity {
		t.Error(" Expected :", expected, " Result: ", p1.Quantity)
	}
}

func TestDisplay(t *testing.T) {

	p1 := products.Product{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 10, Category: "Electronics"}
	//"Products Details are :\nID\t:%d\nName\t:%s\nPrice\t:%.2f\nQuantity\t:%d\nCategory\t:%s", p.ID, p.Name, p.Price, p.Quantity, p.Category
	expected := fmt.Sprintf("Products Details are :\nID\t:%d\nName\t:%s\nPrice\t:%.2f\nQuantity\t:%d\nCategory\t:%s", p1.ID, p1.Name, p1.Price, p1.Quantity, p1.Category)
	str := p1.Display()

	if expected != str {
		t.Error("Expected :", expected, "\nResult :", str)
	}
}
