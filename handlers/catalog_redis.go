package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis-k6-fiber/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogSrv  services.CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(catalogSrv services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{catalogSrv, redisClient}
}

func (h catalogHandlerRedis) GetProduct(c *fiber.Ctx) error {
	key := "handlers::GetProducts"
	// Redis GET
	if responseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}
	// Services
	products, err := h.catalogSrv.GetProduct()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// Redis SET
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}
	fmt.Println("database")
	return c.JSON(response)
}
