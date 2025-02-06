package products

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"category"`
}

func (p *Product) UpdatePrice(newPrice float64) {

	p.Price = newPrice

}

func (p *Product) Sell(quantity int) error {

	if p.Quantity == 0 {
		return errors.New("No further quantities to remove")
	} else {
		p.Quantity = p.Quantity - quantity
	}

	return nil
}

func (p *Product) Restock(quantity int) {

	p.Quantity += quantity

}

func (p *Product) Display() string {
	str := fmt.Sprintf("Products Details are :\nID\t:%d\nName\t:%s\nPrice\t:%.2f\nQuantity\t:%d\nCategory\t:%s", p.ID, p.Name, p.Price, p.Quantity, p.Category)

	return str
}
