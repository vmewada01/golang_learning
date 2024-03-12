package models

import "gorm.io/gorm"

type CustomerStructure struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	UUID    string `gorm:"type:varchar(36);unique_index" json:"uuid"`
	Address string `json:"address"`
	Company string `json:"company"`
}