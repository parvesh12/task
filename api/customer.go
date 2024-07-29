package api

import (
	"task/utils"

	"github.com/gin-gonic/gin"
)

// CustomerAnalysis represents the customer analysis data
type CustomerAnalysis struct {
	TotalCustomers    int     `json:"total_customers"`
	TotalOrders       int     `json:"total_orders"`
	AverageOrderValue float64 `json:"average_order_value"`
}

func GetCustomerAnalysis(c *gin.Context) {

	custcount, err := GetCustomercount()
	if err != nil {
		c.AbortWithError(500, err)
	}

	ordercount, err := GetCustomerOrderCount()
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.JSON(200, CustomerAnalysis{
		TotalCustomers: int(custcount),
		TotalOrders:    int(ordercount),
	})
}

// db query
func GetCustomercount() (count int64, err error) {
	if err := utils.DB.Table("customers").Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func GetCustomerOrderCount() (count int64, err error) {
	if err := utils.DB.Table("orders").Count(&count).Error; err != nil {
		return count, err
	}
	return count, err
}
