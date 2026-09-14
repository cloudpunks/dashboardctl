package kubernetes

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

func buildApiserver() (*dashboardv2.Dashboard, error) {
	builder, err := dashboardv2.NewDashboardBuilder("Kubernetes APIServer").
		Tags([]string{}).
		Build()

	if err != nil {
		return nil, err
	}

	return &builder, nil
}
