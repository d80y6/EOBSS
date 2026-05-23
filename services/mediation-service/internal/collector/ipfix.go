package collector

import (
	"context"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"go.uber.org/zap"
	"net"
)

// IPFIXCollector handles high-throughput ingestion of flow records from network routers
type IPFIXCollector struct {
	port string
}

func (c *IPFIXCollector) StartIngestion(ctx context.Context) error {
	addr := ":" + c.port
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()

	logger.Info("Mediation IPFIX collector started", zap.String("port", c.port))

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Error("Failed to accept connection", zap.Error(err))
			continue
		}
		go c.handleConnection(conn)
	}
}

func (c *IPFIXCollector) handleConnection(conn net.Conn) {
	defer conn.Close()
	// Logic to decode IPFIX templates and data sets
	// and produce UsageEvents to Kafka.
}
