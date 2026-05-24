package service

import (
	"context"
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/repository"
)

type CatalogService struct {
	repo repository.ProductOfferingRepository
}

func NewCatalogService(repo repository.ProductOfferingRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

func (s *CatalogService) CreateOffering(ctx context.Context, offering *domain.ProductOffering) error {
	return s.repo.Create(ctx, offering)
}

func (s *CatalogService) GetOffering(ctx context.Context, id string) (*domain.ProductOffering, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CatalogService) ListOfferings(ctx context.Context) ([]domain.ProductOffering, error) {
	return s.repo.List(ctx)
}
