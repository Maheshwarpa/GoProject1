package inventory

import (
	"errors"
	"fmt"
	Products "module/assignment/products"
)

type Inventory struct {
	Products map[int]Products.Product
}

func (i *Inventory) AddProduct(p Products.Product) {
	_, exists := i.Products[p.ID]

	if exists {
		//fmt.Errorf("Already this product exists in your inventory")

		errors.New("Already this product exists in your inventory")
		return
	} else {
		//Inventory.Products[p.ID] = p
		i.Products[p.ID] = p
	}

}

func (i *Inventory) RemoveProduct(id int) error {

	_, exists := i.Products[id]

	if exists {
		delete(i.Products, id)
		return nil
	} else {
		return errors.New("no product found to remove")
	}

}

func (i *Inventory) FindProductByName(name string) (*Products.Product, error) {
	for _, v := range i.Products {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("product with name %s not found", name)
}

func (i *Inventory) ListByCategory(category string) []Products.Product {
	var mp []Products.Product

	for _, v := range i.Products {
		if v.Category == category {
			mp = append(mp, v)
		}
	}
	return mp
}

func (i *Inventory) TotalValue() float64 {
	sum := 0.0
	for _, v := range i.Products {
		sum += (v.Price * float64(v.Quantity))
		//fmt.Println(sum)
	}
	return sum
}

func (i *Inventory) Save(p Products.Product) error {
	_, exists := i.Products[p.ID]
	if exists {
		return fmt.Errorf(" Product already present")
	} else {
		i.AddProduct(p)
		return nil
	}
}

func (i *Inventory) Zoom() {
	for k, _ := range i.Products {
		delete(i.Products, k)
	}
}

func (i *Inventory) GetByID(id int) (*Products.Product, error) {
	// for _, v := range i.Products {
	// 	if v.ID == id {
	// 		return &v, nil
	// 	}
	// }

	if prod, exists := i.Products[id]; exists {
		return &prod, nil
	}

	return nil, fmt.Errorf("ID not found in Inventory")
}

func (i *Inventory) Delete(id int) error {
	_, exists := i.Products[id]
	if exists {
		return i.RemoveProduct(id)

	}

	return fmt.Errorf("ID not found to perform delete Operations")
}
