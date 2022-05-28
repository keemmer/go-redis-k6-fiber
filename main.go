package main

import (
	"fmt"
	"go-redis-k6-fiber/repositories"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()

	// productRepo := repositories.NewProductRepositoryDB(db)
	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	products, err := productRepo.GetProduct()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(products)

	// app := fiber.New()
	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello world")
	// })

	// app.Listen(":8000")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3307)/test")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "",
		DB:                 0,
		MaxRetries:         2,
		IdleTimeout:        time.Second * 120,
		IdleCheckFrequency: time.Second * 300,
	})
}
