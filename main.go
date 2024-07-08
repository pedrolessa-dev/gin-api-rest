package main

import (
	"github.com/pedrolessa-dev/gin-api-rest/database"
	"github.com/pedrolessa-dev/gin-api-rest/routes"
)

func main() {
	database.DBConnect()
	routes.HandleRequests()
}
