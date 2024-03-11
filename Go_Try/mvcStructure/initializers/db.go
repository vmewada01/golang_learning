package initializers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
    dsn := "root:password@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect to database:", err)
        panic(err)
    }
	DB= db
    // return db, nil
}
