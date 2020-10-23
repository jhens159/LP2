package models

import (
	"gorm.io/gorm"
)


type Docente struct{
	gorm.Model
	Name string
	Paternal string
	Maternal string
	Age string
}

