package main

import (
	"github.com/FlyingQueue/database"
	"github.com/FlyingQueue/routes"
)

func init() {
	database.StartDB()
}
func main() {
	routes.HandleRequests()

}
