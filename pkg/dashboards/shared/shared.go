package shared

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

// BuildFunc builds a single dashboard.
type BuildFunc func() (*dashboardv2.Dashboard, error)
