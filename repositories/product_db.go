package repositories

import (
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db: db}
}

func (r productRepositoryDB) GetProduct() (products []product, err error) {
	// err = r.db.Order("id desc").Limit(50000).Find(&products).Error
	err = r.db.Order("id desc").Limit(50).Find(&products).Error
	return products, err
}
