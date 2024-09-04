package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm="type:varchar(100);not null"`
	Description string `json:"description" gorm="type:text"`
	Completed   bool   `json:"completed" gorm="default:false"`
}
