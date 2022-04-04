package model

import "github.com/jinzhu/gorm"

// User
type Info struct {
	gorm.Model
	Name    string `gorm:"not null;unique;size:32"`
	Timer   int    `gorm:"not null"`
	Min     int
	Max     int
	Boy     int
	Girl    int
	Price   float32
	Npc     int
	IsInd   bool
	IsLimit bool
	Url     string
	Tags    []Tag `gorm:"many2many:Info_Tag"`
}
