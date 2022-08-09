package mgo

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"store/services/models"
)

type ProductRepository interface {
	UpdateProduct() error
	GetProductById(objId string) (*models.Product, error)
	AddProduct(product models.Product) error
	DeleteProduct(objId string) error
}

type productRepository struct {
	collection *mongo.Collection
	redis      *redis.Client
	context    context.Context
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (p *productRepository) UpdateProduct() error {
	return nil
}
func (p *productRepository) GetProductById(obj string) (*models.Product, error) {
	return nil, nil
}
func (p *productRepository) DeleteProduct(obj string) error {
	return nil
}

func (p *productRepository) AddProduct(product models.Product) error {
	return nil
}
