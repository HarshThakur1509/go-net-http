package main

import (
	"test/initializers"
	"test/models"
)

func init() {
	initializers.COnnectDb()
}
func main() {
	Idea := &models.Idea{}
	initializers.DB.AutoMigrate(Idea)

}
