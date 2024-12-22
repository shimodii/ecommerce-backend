package repository

import (
	"github.com/shimodii/ecommerce-backend/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
    Db *gorm.DB
}

var Database DbInstance

func OpenDatabase(){
    db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("databased failed to open!")
    }

    db.AutoMigrate(&entities.Product{})
    db.AutoMigrate(&entities.Cart{})
    db.AutoMigrate(&entities.User{})

    Database = DbInstance{
        Db: db,
    }

}
