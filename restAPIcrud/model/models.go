package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Telnumber int
	IDcard    int
}
