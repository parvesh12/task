package migration

import (
	"task/utils"
	"time"
)

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id         int
	Name       string
	CategoryId int
}

type Customer struct {
	Id      int
	Name    string
	Email   string
	Address string
}

type Order struct {
	Id          int
	CustomerId  int
	OrderDate   time.Time
	TotalAmount string
}

type OrderItem struct {
	Id        int
	ProductId int
	OrderId   int
	Quantity  int
}

type Orders struct {
	ID              int
	ProductID       string
	CustomerID      string
	DateOfSale      string
	Category        string
	Region          string
	ProductName     string
	QuantitySold    int
	UnitPrice       float64
	Discount        float64
	ShippingCost    float64
	PaymentMethod   string
	CustomerName    string
	CustomerEmail   string
	CustomerAddress string
}

func Migration() {
	err := utils.DB.AutoMigrate(
		&Category{},
		&Product{},
		&Customer{},
		&Order{},
		&OrderItem{},
		&Orders{},
	)

	if err != nil {
		panic(err)
	}
}
