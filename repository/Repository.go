package repository

import (
	"github.com/shimodii/ecommerce-backend/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDatabase(){
    dsn := "host=localhost user=postgres password=pgpass dbname=ecommerce port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("database is fucked up dude")
    }
    
    println("database connected")

    db.AutoMigrate(&entities.User{}, &entities.Product{}, &entities.Cart{})
}
