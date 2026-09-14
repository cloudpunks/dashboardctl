package command

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/cloudpunks/dashboardctl/pkg/dashboards"
)

var (
	outputCmd = &cobra.Command{
		Use:  "output <category>/<dashboard>",
		Run:  outputAction,
		Args: cobra.ArbitraryArgs,
	}
)

func init() {
	rootCmd.AddCommand(outputCmd)
}

func outputAction(_ *cobra.Command, args []string) {
	category, name, err := dashboardNames(args)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to resolve requested dashboard")

		os.Exit(1)
	}

	result, err := dashboards.Build(category, name)

	if err != nil {
		log.Error().Err(err).Str("dashboard", name).Msg("Failed to build dashboard")
		os.Exit(1)
	}

	raw, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		log.Error().Err(err).Str("dashboard", name).Msg("Failed to encode dashboard")
		os.Exit(1)
	}

	fmt.Println(string(raw))
}

func dashboardNames(requested []string) (string, string, error) {
	if len(requested) != 1 {
		return "", "", fmt.Errorf("exactly one dashboard must be specified")
	}

	segments := strings.SplitN(
		requested[0],
		"/",
		2,
	)

	if len(segments) != 2 {
		return "", "", fmt.Errorf("dashboard must be specified as <category>/<dashboard>")
	}

	category := segments[0]
	name := segments[1]

	known := make(map[string]bool)

	for _, n := range dashboards.Names(category) {
		known[n] = true
	}

	if !known[name] {
		return "", "", fmt.Errorf("unknown dashboard %q, run 'dashboardctl list' to see the available ones", name)
	}

	return category, name, nil
}
