package kubernetes

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

func buildNamespaces() (*dashboardv2.Dashboard, error) {
	builder, err := dashboardv2.NewDashboardBuilder("Kubernetes Namespaces").
		Tags([]string{}).
		Build()

	if err != nil {
		return nil, err
	}

	return &builder, nil
}
