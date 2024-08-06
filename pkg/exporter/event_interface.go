package exporter

import (
	"context"
	"github.com/kubesphere/kube-events/pkg/config"
)

type EventSource interface {
	ReloadConfig(c *config.ExporterConfig)

	Run(ctx context.Context) error
}
