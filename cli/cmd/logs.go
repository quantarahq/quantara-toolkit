package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/quantarahq/quantara-toolkit/cli/internal/apiclient"
	"github.com/spf13/cobra"
)

type logEvent struct {
	at   time.Time
	line string
}

var logsCmd = &cobra.Command{
	Use:   "logs <project-id>",
	Short: "Show deployment and contract activity for a project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := args[0]
		client := apiclient.New(apiURL)
		ctx := cmd.Context()

		deployments, err := client.ListDeployments(ctx, projectID)
		if err != nil {
			return fmt.Errorf("fetching deployments: %w", err)
		}
		contracts, err := client.ListContracts(ctx, projectID)
		if err != nil {
			return fmt.Errorf("fetching contracts: %w", err)
		}

		events := make([]logEvent, 0, len(deployments)+len(contracts))
		for _, d := range deployments {
			events = append(events, logEvent{
				at: d.CreatedAt,
				line: fmt.Sprintf("[deployment] %s  contract=%s  status=%s",
					d.DeploymentID, d.ContractName, d.Status),
			})
		}
		for _, c := range contracts {
			events = append(events, logEvent{
				at: c.Timestamp,
				line: fmt.Sprintf("[contract]   %s  address=%s  hash=%s",
					c.DeploymentID, c.ContractAddress, c.DeploymentHash),
			})
		}

		if len(events) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No activity found for this project.")
			return nil
		}

		sort.Slice(events, func(i, j int) bool { return events[i].at.Before(events[j].at) })

		out := cmd.OutOrStdout()
		for _, e := range events {
			fmt.Fprintf(out, "%s  %s\n", e.at.Format("2006-01-02 15:04:05"), e.line)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
