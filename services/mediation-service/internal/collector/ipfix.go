package collector

import (
	"context"
	"fmt"
	"net"
	"time"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"go.uber.org/zap"
)

// IPFIXCollector listens for NetFlow/IPFIX binary packets from routers
type IPFIXCollector struct {
	Port int
}

func (c *IPFIXCollector) Start(ctx context.Context, dataChan chan<- map[string]interface{}) error {
	addr := fmt.Sprintf(":%d", c.Port)
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	logger.Info("IPFIX Collector started", zap.Int("port", c.Port))

	buf := make([]byte, 2048)
	for {
		select {
		case <-ctx.Done():
			logger.Info("IPFIX Collector shutting down")
			return nil
		default:
			// Set a deadline for ReadFrom to prevent indefinite blocking and allow ctx.Done check
			_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
			n, remoteAddr, err := conn.ReadFrom(buf)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue // Normal timeout, check ctx again
				}
				continue
			}

			logger.Info("Received IPFIX packet",
				zap.Int("bytes", n),
				zap.String("remote", remoteAddr.String()),
			)

			// Mock parsed data for transformer pipeline
			dataChan <- map[string]interface{}{
				"flow_id":   fmt.Sprintf("%s-%d", remoteAddr.String(), n),
				"bytes":     float64(n),
				"router_id": remoteAddr.String(),
				"src_ip":    "10.0.0.1",
				"service_id": "SVC-MOCK-99", // Providing required service_id
			}
		}
	}
}
