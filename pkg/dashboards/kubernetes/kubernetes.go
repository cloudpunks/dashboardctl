package kubernetes

import (
	"fmt"
	"sort"

	"github.com/cloudpunks/dashboardctl/pkg/dashboards/shared"
	"github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

// Registry maps every known dashboard name to its builder.
var Registry = map[string]shared.BuildFunc{
	"apiserver":    buildApiserver,
	"applications": buildApplications,
	"cluster":      buildCluster,
	"coredns":      buildCoreDNS,
	"deployments":  buildDeployments,
	"global":       buildGlobal,
	"namespaces":   buildNamespaces,
	"nodes":        buildNodes,
	"pods":         buildPods,
}

// Names returns the sorted list of all known dashboard names.
func Names() []string {
	names := make([]string, 0, len(Registry))

	for name := range Registry {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Build builds the dashboard registered under the given name.
func Build(name string) (*dashboardv2.Dashboard, error) {
	build, ok := Registry[name]

	if !ok {
		return nil, fmt.Errorf("unknown dashboard %q, run 'dashboardctl list' to see the available ones", name)
	}

	return build()
}
