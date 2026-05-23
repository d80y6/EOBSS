package repository

import (
	"context"
	"database/sql"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
)

type ClickHouseCDRRepository struct {
	db *sql.DB
}

func NewClickHouseCDRRepository(db *sql.DB) *ClickHouseCDRRepository {
	return &ClickHouseCDRRepository{db: db}
}

func (r *ClickHouseCDRRepository) InsertCDR(ctx context.Context, record *domain.UsageRecord) error {
	query := "INSERT INTO cdr (id, usage_type, quantity, amount, timestamp, service_id) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query,
		record.ID,
		record.UsageType,
		record.UsageQuantity,
		record.RatedAmount,
		record.Timestamp,
		record.ServiceRef,
	)
	return err
}
