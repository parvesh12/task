package api

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"task/migration"
	"task/utils"

	"github.com/gin-gonic/gin"
)

type JsonMessage struct {
	Message string
}

func LoadCSV(c *gin.Context) {

	file, err := os.Open("product.csv")
	if err != nil {
		c.AbortWithError(500, err)
	}

	defer file.Close()
	csvreader := csv.NewReader(file)
	records, err := csvreader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	for i, record := range records {
		if i >= 1 {
			order := migration.Orders{
				ID:            atoi(record[0]),
				ProductID:     record[1],
				CustomerID:    record[2],
				ProductName:   record[3],
				Category:      record[5],
				DateOfSale:    record[7],
				QuantitySold:  atoi(record[7]),
				UnitPrice:     atof(record[8]),
				Discount:      atof(record[9]),
				ShippingCost:  atof(record[10]),
				PaymentMethod: record[11],
			}
			err := utils.DB.Exec(`INSERT INTO orders (product_id, customer_id, date_of_sale, quantity_sold, unit_price, discount, shipping_cost, payment_method) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, order.ProductID, order.CustomerID, order.DateOfSale, order.QuantitySold, order.UnitPrice, order.Discount, order.ShippingCost, order.PaymentMethod)
			if err.Error != nil {
				fmt.Println(err)
				c.AbortWithError(500, err.Error)
			}
		}

	}

	c.JSON(200, JsonMessage{Message: "Csv load successfully"})

}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
