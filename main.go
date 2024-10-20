package main

import (
	"github.com/HarshThakur1509/go-net-http/api"
	"github.com/HarshThakur1509/go-net-http/initializers"
)

func init() {
	initializers.COnnectDb()
}

func main() {
	server := api.NewAPIServer(":3000")
	server.Run()
}
