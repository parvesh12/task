package api

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"task/migration"
	"task/utils"

	"github.com/gin-gonic/gin"
)

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
	for _, record := range records {

		order := migration.Orders{
			ProductID:     atoi(record[1]),
			CustomerID:    atoi(record[2]),
			DateOfSale:    record[3],
			QuantitySold:  atoi(record[4]),
			UnitPrice:     atof(record[5]),
			Discount:      atof(record[6]),
			ShippingCost:  atof(record[7]),
			PaymentMethod: record[8],
		}

		err := utils.DB.Exec(`INSERT INTO orders (product_id, customer_id, date_of_sale, quantity_sold, unit_price, discount, shipping_cost, payment_method) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, order.ProductID, order.CustomerID, order.DateOfSale, order.QuantitySold, order.UnitPrice, order.Discount, order.ShippingCost, order.PaymentMethod)
		if err != nil {
			c.AbortWithError(500, err.Error)
		}
	}

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
