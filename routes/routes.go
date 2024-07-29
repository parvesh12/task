package routes

import (
	"task/api"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	r := gin.Default()
	r.POST("/loadcsv", api.LoadCSV)
	r.POST("/customeranalysis", api.GetCustomerAnalysis)
	return r
}
