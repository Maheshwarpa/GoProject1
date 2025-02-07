package main

import (
	"fmt"
	api "module/assignment/API"
	"module/assignment/Database"
	inventory "module/assignment/Inventory"
	products "module/assignment/products"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// var p inventory.Inventory
var url api.ExternalAPI

func getAllItem(b *gin.Context, inv *inventory.Inventory) {

	b.IndentedJSON(http.StatusOK, inv.Products)

}

func getItemByCategory(b *gin.Context, inv *inventory.Inventory) {

	/*idCateg := b.Request.URL.String()

	str:= strings.Split(idCateg, "/")*/

	id := b.Param("category")

	cat := inv.ListByCategory(id)

	if cat == nil {
		b.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category found"})
		return
	}

	b.JSON(http.StatusOK, cat)

}

func getItemByName(b *gin.Context, inv *inventory.Inventory) {
	//idName := b.Param("name")
	url.BaseURL = b.Request.URL.String()
	fmt.Println(url.BaseURL)
	str := strings.Split(url.BaseURL, "/")
	fmt.Println(str)
	key := str[len(str)-1]
	fmt.Println(key)
	product, err := inv.FindProductByName(key)
	if err != nil {
		b.JSON(http.StatusNotFound, gin.H{"error": "Product not found with given name"})
		return
	}

	b.IndentedJSON(http.StatusOK, *product)

}

func getItem(b *gin.Context, inv *inventory.Inventory) {

	url.BaseURL = b.Request.URL.String()
	fmt.Println(url.BaseURL)
	str := strings.Split(url.BaseURL, "/")
	key := str[len(str)-1]
	fmt.Println(key)
	//idParm := b.Param("id")
	// idName := b.Param("name")
	// fmt.Println(idName)
	id, err := strconv.Atoi(key)
	if err != nil {
		b.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	//b.IndentedJSON(http.StatusAccepted, id)
	product, err := inv.GetByID(id)
	if err != nil {
		fmt.Println(err)
		b.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	fmt.Println(*product)
	b.IndentedJSON(http.StatusOK, *product)

}

func saveItem(b *gin.Context, inv *inventory.Inventory) {

	p1 := products.Product{ID: 30, Name: "Leather Wallet", Price: 39.99, Quantity: 60, Category: "Accessories"}

	if err := b.ShouldBindJSON(&p1); err != nil {
		b.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := inv.Save(p1)
	if err != nil {
		fmt.Errorf("error is:%s", err)
		b.JSON(http.StatusBadRequest, gin.H{"error": "Already Product Present"})
		return
	}

	b.JSON(http.StatusCreated, gin.H{"message": "Product added successfully", "product": p1})

}

func deleteItem(b *gin.Context, inv *inventory.Inventory) {
	url.BaseURL = b.Request.URL.String()
	fmt.Println(url.BaseURL)
	str := strings.Split(url.BaseURL, "/")
	key := str[len(str)-1]
	fmt.Println(key)

	//idParm := b.Param("id")

	id, err := strconv.Atoi(key)
	if err != nil {
		b.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	err1 := inv.Delete(id)

	if err1 != nil {
		fmt.Println(err1)
		b.JSON(http.StatusNotFound, gin.H{"error": "Delete option not found"})
		return
	}

	b.JSON(http.StatusOK, gin.H{"Message": "Successfully Deleted the product"})

}

func main() {
	inv := &inventory.Inventory{
		Products: make(map[int]products.Product), // Initialize the Products map
	}

	Database.LoadData(inv)
	fmt.Println(inv.Products)
	/*	var id int
		fmt.Scan(&id)
		prod, err := inv.GetByID(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*prod)
		fmt.Scan(&id)
		err1 := inv.Delete(id)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		fmt.Println(inv.Products)
		fmt.Println("Hello")*/
	//var p inventory.Inventory
	router := gin.Default()
	router.GET("/Items", func(b *gin.Context) { getAllItem(b, inv) })
	router.GET("/Items/:id", func(b *gin.Context) { getItem(b, inv) })
	router.GET("/Items/name/:name", func(b *gin.Context) { getItemByName(b, inv) })
	router.GET("/Items/category/:category", func(b *gin.Context) { getItemByCategory(b, inv) })
	router.POST("/Items", func(b *gin.Context) { saveItem(b, inv) })
	router.DELETE("/Items/:id", func(b *gin.Context) { deleteItem(b, inv) })
	router.Run("localhost:8080")

}
