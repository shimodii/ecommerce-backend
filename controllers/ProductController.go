package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/ecommerce-backend/entities"
	"github.com/shimodii/ecommerce-backend/repository"
)

func GetAllProducts(c *fiber.Ctx) error {
    var products []entities.Product

	repository.Database.Db.Find(&products)

	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product entities.Product

	err := repository.Database.Db.Find(&product, "id = ?", id)
	if err != nil {
		panic(err)
	}

	return c.JSON(product)
}

func AddProduct(c *fiber.Ctx) error {
	var product entities.Product

	err := c.BodyParser(&product)
	if err != nil {
		panic(err)
	}

	repository.Database.Db.Create(&product)

	return c.Status(200).JSON("OK")
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product entities.Product

	err := repository.Database.Db.Find(&product, "id = ?", id)
	if err != nil {
		return c.Status(404).JSON("product not found")
	}

	err2 := repository.Database.Db.Delete(&product).Error
	if err2 != nil {
		return c.Status(400).JSON("product not deleted")
	}

	return c.Status(200).JSON("OK")
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
    
    var product entities.Product

    err := repository.Database.Db.Find(&product, "id = ?", id)
    if err != nil {
        c.Status(404).JSON("product not found")
    }

    type newProductData struct {
        Name string `json:"name"`
        Description string `json:"description"`
		Price float64 `json:"price"`
		Stock int `json:"stock"`
    }
    var newProduct newProductData

    err2 := c.BodyParser(&newProduct) 
	if err2 != nil {
        c.Status(400).JSON("error")
    }

    product.Name = newProduct.Name
    product.Description = newProduct.Description
    product.Price = newProduct.Price
    product.Stock = newProduct.Stock

    repository.Database.Db.Save(&product)
    
    return c.Status(200).JSON(product)
}
