package main

import (
	"test/api"
	"test/initializers"
)

func init() {
	initializers.COnnectDb()
}

func main() {
	server := api.NewAPIServer(":3000")
	server.Run()
}
