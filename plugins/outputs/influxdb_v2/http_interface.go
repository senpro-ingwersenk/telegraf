package influxdb_v2

import (
	"context"
	"net/url"

	"github.com/influxdata/telegraf"
)

type httpExecutor interface {
	Init() error
	GetUrl() *url.URL
	Write(ctx context.Context, metrics []telegraf.Metric) error
	Close()
}
