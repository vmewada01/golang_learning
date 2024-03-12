package initializers

import (
	"fmt"
	"log"
	"mvcStructure/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
    dsn := "root:password@tcp(127.0.0.1:3306)/learningGo?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect to database:", err)
        panic(err)
    } else {
		log.Println("Database connection established")
	}


    err = db.AutoMigrate(&models.CustomerStructure{}, &models.UserStructure{}, &models.TodosModel{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Auto migration completed")
	}

	DB= db
    // return db, nil
}
