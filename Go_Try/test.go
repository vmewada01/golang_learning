package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       int64
	Birthday  time.Time
	UUID      string `gorm:"default:uuid_generate_v3()"`
	Confirmed bool
	Address   Address
	MemberShip string // Define MemberShip field
}

type Address struct {
	gorm.Model
	// Define fields for Address struct here
}

type Person struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age"`
	Email  string `json:"email,omitempty"`
	Active bool   `json:"active,omitempty"`
}

func main() {


    p1 := Person{Name: "Alice", Age: 30, Email: "", Active: true}
	p2 := Person{Name: "", Age: 0, Email: "bob@example.com", Active: false}

	// Marshal structs to JSON
	jsonData1, _ := json.Marshal(p1)
	jsonData2, _ := json.Marshal(p2)

	// Print JSON data
	fmt.Println(string(jsonData1)) // {"name":"Alice","age":30,"active":true}
	fmt.Println(string(jsonData2)) // {"email":"bob@example.com"}

	fmt.Println("vishal p1 print",p1)
	fmt.Println("vishal p2 print", p2)


	// dsn := "root:password@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Create user
	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// result := db.Create(&user)

    // fmt.Println("result", result)
	// // GORM supports create from map:
	// db.Model(&User{}).Create(map[string]interface{}{
	// 	"Name": "jinzhu", "Age": 18,
	// })

	// // Create With Associations
	// db.Create(&User{
	// 	Name:      "jinzhu",
	// 	Confirmed: true,
	// 	Address:   Address{},
	// })
}

// Hooks on Create Actions
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	if !u.IsValid() {
		err = errors.New("can't save invalid data")
	}
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("role", "admin")
	}
	if !u.IsValid() {
		return errors.New("rollback invalid user")
	}
	return
}

// IsValid method definition
func (u *User) IsValid() bool {
	// Implement your validation logic here
	return true // Change this based on your validation criteria
}

// Hooks on Update Actions
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// if u.Address() {
	// 	err = errors.New("read only user")
	// }
	 return
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	if u.Confirmed {
		tx.Model(&Address{}).Where("user_id = ?", u.ID).Update("verified", true)
	}
	return
}

// Hooks on Delete Actions
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	if u.Confirmed {
		tx.Model(&Address{}).Where("user_id = ?", u.ID).Update("invalid", false)
	}
	return
}

// Hooks on Read Actions
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.MemberShip == "" {
		u.MemberShip = "user"
	}
	return
}
