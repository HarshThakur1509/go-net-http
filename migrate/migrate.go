package main

import (
	"github.com/HarshThakur1509/go-net-http/initializers"
	"github.com/HarshThakur1509/go-net-http/models"
)

func init() {
	initializers.COnnectDb()
}
func main() {
	Idea := &models.Idea{}
	initializers.DB.AutoMigrate(Idea)

}
