package main

import (
	"go-redis-k6-fiber/handlers"
	"go-redis-k6-fiber/repositories"
	"go-redis-k6-fiber/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Repository
	// redisClient := initRedis()
	// db := initDatabase()
	// productRepo := repositories.NewProductRepositoryDB(db)
	// =========================================================================
	// db := initDatabase()
	// redisClient := initRedis()
	// productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	// products, err := productRepo.GetProduct()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(products)

	// Services
	// =========================================================================
	// db := initDatabase()
	// productRepo := repositories.NewProductRepositoryDB(db)
	// productService := services.NewCatalogService(productRepo)
	// products, err := productService.GetProduct()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(products)
	// =========================================================================
	// db := initDatabase()
	// redisClient := initRedis()
	// productRepo := repositories.NewProductRepositoryDB(db)
	// productService := services.NewCatalogServiceRedis(productRepo,redisClient)
	// products, err := productService.GetProduct()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(products)

	// Handler
	// =========================================================================
	// db := initDatabase()
	// redisClient := initRedis()
	// productRepo := repositories.NewProductRepositoryDB(db)
	// productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	// productHandler := handlers.NewCatalogHandler(productService)
	// =========================================================================
	db := initDatabase()
	redisClient := initRedis()
	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogService(productRepo)
	productHandler := handlers.NewCatalogHandlerRedis(productService, redisClient)

	app := fiber.New()
	app.Get("/products", productHandler.GetProduct)
	app.Listen(":8000")
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
