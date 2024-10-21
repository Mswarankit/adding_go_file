package main

import (
	"fmt"
	"strings"
)

type Customer struct {
	Name           string
	Phone          string
	PhoneId        string
	Item           map[string]float64
	ItemsPurchased string
	BillPayed      string
}

func (c *Customer) CustomerName(Name string) *Customer {
	c.Name = Name
	return c
}

func (c *Customer) CustomerPhone(Phone string) *Customer {
	c.Phone = Phone
	return c
}

func (c *Customer) CustomerItems(Item map[string]float64) *Customer {
	if c.Item == nil {
		c.Item = make(map[string]float64)
	}
	for key, value := range Item {
		c.Item[key] = value
	}
	var itemsList []string
	for key, value := range c.Item {
		itemsList = append(itemsList, fmt.Sprintf("%s: ₹%.2f", key, value))
	}
	c.ItemsPurchased = strings.Join(itemsList, ", ")
	return c
}

func (c *Customer) CustomerBillPayment(Item map[string]float64) *Customer {
	var sum float64
	for _, value := range Item {
		sum += value
	}
	taxRate := 0.12
	totalWithTax := sum + (sum * taxRate)
	c.BillPayed = fmt.Sprintf("%s %.2f", "\u20B9", totalWithTax)
	return c
}

func (c *Customer) CustomerFairPoint(PhoneId string) *Customer {
	c.PhoneId = fmt.Sprintf("₹%s", PhoneId)
	return c
}

func (c *Customer) PrintAll() {
	fmt.Printf("Dear %s, Thanks for Shopping with us.\nHere are the items: %s.\nHere is the total price: %s.\nYour phone number: %s has added %s for next shopping.\n",
		c.Name, c.ItemsPurchased, c.BillPayed, c.Phone, c.PhoneId)
}

func main() {
	customer := Customer{}
	items := map[string]float64{
		"Tool1": 14.5,
		"Tool2": 17.8,
		"Tool3": 19.9,
		"Tool4": 24.5,
	}
	customer.CustomerName("John Doe").
		CustomerPhone("2080008080").CustomerItems(items).CustomerBillPayment(items).CustomerFairPoint("10").PrintAll()
}
