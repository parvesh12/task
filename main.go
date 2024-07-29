package main

import (
	"os"
	"task/migration"
	"task/routes"
)

func main() {

	migration.Migration() //table migration

	r := routes.SetupRoutes() //  intialize routes

	r.Run(":" + os.Getenv("PORT"))

}
