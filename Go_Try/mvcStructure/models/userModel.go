package models

type UserStructure struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Status  bool   `json:"status"`
	Age     int    `json:"age"`
	UUID    string `gorm:"default:uuid_generate_v3()"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}