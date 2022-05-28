package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis-k6-fiber/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{productRepo, redisClient}
}
func (s catalogServiceRedis) GetProduct() (products []Product, err error) {

	key := "services::GetProducts"
	// Redis GET
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}
	// Repository
	productsDB, err := s.productRepo.GetProduct()
	if err != nil {
		return nil, err
	}
	for _, p := range productsDB {
		products = append(products, Product{ID: int(p.ID), Name: p.Name, Quantity: p.Quantity})
	}
	// Redis SET
	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}
	fmt.Println("database")
	return products, nil
}
