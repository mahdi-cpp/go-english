package models

import (
	"fmt"
	"github.com/mahdi-cpp/go-english/config"
)

type User struct {
	ID    uint
	Name  string
	Email string
	Info  map[string]interface{} `gorm:"type:json;serializer:json"`
}

func CreatUsers() {
	// Create a new user record with JSONB data
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Info:  map[string]interface{}{"age": 30, "city": "tabrize", "width": 100, "height": 2000},
	}
	config.DB.Create(&user)
}

func QueryUsers() {
	var users []User

	config.DB.Where("info->>'width' = ?", 11).Find(&users)

	fmt.Println("query result: ")
	for _, u := range users {
		var city = u.Info["city"]
		fmt.Println(city)
	}

	fmt.Println("-----------------------")
}
