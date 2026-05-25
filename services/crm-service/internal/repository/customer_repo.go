package repository

import (
	"context"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/service"
	"gorm.io/gorm"
)

type gormCustomerRepo struct {
	db *gorm.DB
}

func NewGormCustomerRepo(db *gorm.DB) service.CustomerRepository {
	return &gormCustomerRepo{db: db}
}

func (r *gormCustomerRepo) Create(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

func (r *gormCustomerRepo) GetByID(ctx context.Context, id string) (*domain.Customer, error) {
	var customer domain.Customer
	if err := r.db.WithContext(ctx).First(&customer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *gormCustomerRepo) Update(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}
