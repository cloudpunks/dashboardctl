package general

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

func buildNode() (*dashboardv2.Dashboard, error) {
	builder, err := dashboardv2.NewDashboardBuilder("Node Exporter").
		Tags([]string{}).
		Build()

	if err != nil {
		return nil, err
	}

	return &builder, nil
}
