package models

import "gorm.io/gorm"

type TodosModel struct {
	gorm.Model
	TodoName    string `json:"todoName"`
	Value       string `json:"value"`
	UUID        string `gorm:"type:varchar(36);unique_index" json:"uuid"`
	Status      string `json:"status"`
	Description string `json:"description"`}