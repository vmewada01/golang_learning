package models

import "gorm.io/gorm"

type UserStructure struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Status  bool   `json:"status"`
	Age     int    `json:"age"`
	UUID    string `gorm:"type:varchar(36);unique_index" json:"uuid"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}