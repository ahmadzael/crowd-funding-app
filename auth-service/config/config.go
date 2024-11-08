package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := "username:password@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }
    DB = db
    fmt.Println("Database connected!")
}
