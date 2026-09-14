package command

import (
	"fmt"
	"strings"

	"github.com/cloudpunks/dashboardctl/pkg/dashboards"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List the available dashboards",
		Run:   listAction,
		Args:  cobra.NoArgs,
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}

func listAction(_ *cobra.Command, _ []string) {
	for _, category := range dashboards.Categories() {
		for _, name := range dashboards.Names(category) {
			fmt.Println(strings.Join(
				[]string{
					category,
					name,
				},
				"/",
			))
		}
	}
}
