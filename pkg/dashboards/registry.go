// Package dashboards is a registry of all the Grafana dashboards known to
// dashboardctl, each implemented as its own Go package built with the
// Grafana foundation SDK.
package dashboards

import (
	"fmt"
	"sort"

	"github.com/cloudpunks/dashboardctl/pkg/dashboards/general"
	"github.com/cloudpunks/dashboardctl/pkg/dashboards/kubernetes"
	"github.com/cloudpunks/dashboardctl/pkg/dashboards/shared"
	"github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

// Registry maps every known dashboard name to its builder.
var Registry = map[string]map[string]shared.BuildFunc{
	"general":    general.Registry,
	"kubernetes": kubernetes.Registry,
}

// Categories returns the sorted list of all known dashboard categories.
func Categories() []string {
	names := make([]string, 0, len(Registry))

	for name := range Registry {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Names returns the sorted list of all known dashboard names.
func Names(category string) []string {
	if _, ok := Registry[category]; !ok {
		return []string{}
	}

	names := make([]string, 0, len(Registry[category]))

	for name := range Registry[category] {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Build builds the dashboard registered under the given name.
func Build(category, name string) (*dashboardv2.Dashboard, error) {
	builds, ok := Registry[category]

	if !ok {
		return nil, fmt.Errorf("unknown dashboard category %q, run 'dashboardctl list' to see the available ones", category)
	}

	build, ok := builds[name]

	if !ok {
		return nil, fmt.Errorf("unknown dashboard %q/%q, run 'dashboardctl list' to see the available ones", category, name)
	}

	return build()
}
