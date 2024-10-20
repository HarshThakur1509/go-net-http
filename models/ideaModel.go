package models

import "gorm.io/gorm"

type Idea struct {
	gorm.Model
	Title string
}
