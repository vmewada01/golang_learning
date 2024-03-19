package models

import (
	"github.com/vmewada01/todo/config"
	"gorm.io/gorm"
)


var db *gorm.DB 

type Todo struct{
	gorm.Model 
	Title string `gorm:"" json:"title"`
    Description string `json:"description"`
	Content string `gorm:"" json:"content"`
	Status bool `json:"status"`
	
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}


func (b *Todo) CreateTodo()*Todo{
	db.Create(&b)
	return b
}

func GetAllTodo() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos

}

func GetTodoByID(ID int64) (*Todo, error) {
    var todo Todo
    result := db.Where("ID = ?", ID).First(&todo)
    if result.Error != nil {
        return nil, result.Error
    }
    return &todo, nil
}



func DeleteTodo(ID int64) error {
    var todo Todo
    result := db.Where("ID = ?", ID).Delete(&todo)
    if result.Error != nil {
        return result.Error
    }
    return nil
}
