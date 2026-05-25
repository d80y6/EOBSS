package service

import (
	"context"
	"time"
	"github.com/telcoflow/telcoflow/services/incident-service/internal/domain"
)

type IncidentService struct {
	// repo repository.TicketRepository
}

func (s *IncidentService) CreateTicket(ctx context.Context, ticket *domain.TroubleTicket) error {
	ticket.ID = "TKT-" + time.Now().Format("20060102150405")
	ticket.CreationDate = time.Now()
	ticket.Status = "Open"
	return nil
}
