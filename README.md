# GoProject1

## Inventory Management System

This project is a simple Inventory Management System written in Go. It consists of multiple packages that handle product storage, inventory management, database operations, API integrations, and testing.


### Packages

#### 1. API

Description:
The api package is responsible for handling interactions with external APIs. It defines a structure to manage external API configurations.

#### Usage:

- Used to store and configure the base URL for external API communication.
- Can be extended to include methods for making API requests.

#### 2. Database

Description:
The database package provides an interface for storing, retrieving, and deleting products in the system. It also contains a function to load sample data into the inventory.

#### Functions:

- LoadData(p *inventory.Inventory): Loads sample product data into the inventory.
- Implements ProductStorage interface for database operations.

#### Usage:

- Facilitates storing and retrieving products from the database.
- Can be integrated with a real database in the future.


#### 3. Inventory

Description:
The inventory package manages the collection of products, allowing addition, removal, searching, and category-wise listing.

#### Functions:

- AddProduct(p Products.Product): Adds a product to inventory.
- RemoveProduct(id int): Removes a product from inventory by ID.
- FindProductByName(name string): Searches for a product by name.
- ListByCategory(category string): Lists all products under a given category.
- TotalValue(): Calculates the total value of all inventory products.
- GetByID(id int): Fetches a product by ID.
- Delete(id int): Deletes a product from inventory.

#### Usage:

- Central module managing inventory operations.
- Supports CRUD operations for product management.

#### 4. Products

Description:
The products package defines the Product structure and provides methods to modify product attributes.

#### Functions:

- UpdatePrice(newPrice float64): Updates product price.
- Sell(quantity int) error: Decreases product quantity when sold.
- Restock(quantity int): Increases product quantity.
- Display() string: Displays product details.

#### Usage:

- Used as a base structure for all products.
- Facilitates modification and retrieval of product details.

#### 5. Test

Description:
The test package contains unit tests to validate the functionality of the inventory system.

#### Test Cases:

- TestAddProduct(): Ensures products are correctly added to inventory.
- TestRemoveProduct(): Verifies product removal from inventory.
- TestFindProductByName(): Checks the product search function.
- TestTotalValue(): Validates total inventory value calculation.

#### Usage:

- Ensures inventory management operations function correctly.
- Helps maintain code quality and reliability.


#### 6. Main

Description:
The main package is the entry point of the Inventory Management System. It initializes the inventory, loads data, and sets up API routes using the Gin framework.

#### Functions:

- getAllItem(): Retrieves all products.
- getItemByCategory(): Fetches products based on category.
- getItemByName(): Searches for a product by name.
- getItem(): Retrieves a product by ID.
- saveItem(): Adds a new product to inventory.
- deleteItem(): Removes a product from inventory.

#### Usage:

- Runs a REST API server using Gin.
- Handles product-related operations through HTTP requests.


### Installation & Usage

#### 1. Clone the repository:

- git clone https://github.com/Maheshwarpa/GoProject1.git

#### 2. Navigate to the project directory:

- cd GoProject1

#### 3. Run tests:

- go test ./test

#### 4. Start the server:

- go run main.go

- Use the API endpoints to manage inventory.

### API Testing with Postman and Command Prompt

##### Postman

Open Postman and create a new request.

Use the following API endpoints:

- Get all items: GET http://localhost:8080/Items

- Get item by ID: GET http://localhost:8080/Items/{id}

- Get item by name: GET http://localhost:8080/Items/name/{name}

- Get item by category: GET http://localhost:8080/Items/category/{category}

- Add an item: POST http://localhost:8080/Items

- Delete an item: DELETE http://localhost:8080/Items/{id}

- For POST requests, go to the Body tab and select raw with JSON format, then enter the product data:

```json   
  {
    "id": 10,
    "name": "New Product",
    "price": 50.99,
    "quantity": 20,
    "category": "Electronics"
  }
``` 

##### Command Prompt

- Use curl commands to test API calls:

 ```curl  
  curl -X GET http://localhost:8080/Items
  curl -X GET http://localhost:8080/Items/{id}
  curl -X POST http://localhost:8080/Items -H "Content-Type: application/json" -d '{"id": 10, "name": "New Product", "price": 50.99, "quantity": 20, "category": "Electronics"}'
  curl -X DELETE HTTP://localhost:8080/Items/{id}

```
#### Future Improvements

- Connect database package to a real database (e.g., PostgreSQL, MySQL).
- Implement API communication in api package.
- Enhance test coverage with more edge cases.
- Add CLI or Web UI for better usability.
