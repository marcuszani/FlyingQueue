package main

import (
	"github.com/FlyingQueue/config"
	"github.com/FlyingQueue/routes"
)

func main() {
	config.Configuration()
	routes.HandleRequests()

}
