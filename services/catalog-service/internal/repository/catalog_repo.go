package repository

import (
	"context"
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/domain"
	"gorm.io/gorm"
)

type ProductOfferingRepository interface {
	Create(ctx context.Context, offering *domain.ProductOffering) error
	GetByID(ctx context.Context, id string) (*domain.ProductOffering, error)
	List(ctx context.Context) ([]domain.ProductOffering, error)
}

type gormCatalogRepo struct {
	db *gorm.DB
}

func NewGormCatalogRepo(db *gorm.DB) ProductOfferingRepository {
	return &gormCatalogRepo{db: db}
}

func (r *gormCatalogRepo) Create(ctx context.Context, offering *domain.ProductOffering) error {
	return r.db.WithContext(ctx).Create(offering).Error
}

func (r *gormCatalogRepo) GetByID(ctx context.Context, id string) (*domain.ProductOffering, error) {
	var offering domain.ProductOffering
	if err := r.db.WithContext(ctx).First(&offering, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &offering, nil
}

func (r *gormCatalogRepo) List(ctx context.Context) ([]domain.ProductOffering, error) {
	var offerings []domain.ProductOffering
	if err := r.db.WithContext(ctx).Find(&offerings).Error; err != nil {
		return nil, err
	}
	return offerings, nil
}
