package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/quantarahq/quantara-toolkit/cli/internal/apiclient"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List and inspect Quantara projects",
}

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := apiclient.New(apiURL)
		projects, err := client.ListProjects(cmd.Context())
		if err != nil {
			return err
		}

		if len(projects) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No projects found.")
			return nil
		}

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 2, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tCREATED")
		for _, p := range projects {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", p.ID, p.Name, p.Description, p.CreatedAt.Format("2006-01-02 15:04:05"))
		}
		return w.Flush()
	},
}

var projectsGetCmd = &cobra.Command{
	Use:   "get <project-id>",
	Short: "Show a single project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := apiclient.New(apiURL)
		project, err := client.GetProject(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		out := cmd.OutOrStdout()
		fmt.Fprintf(out, "ID:          %d\n", project.ID)
		fmt.Fprintf(out, "Name:        %s\n", project.Name)
		fmt.Fprintf(out, "Description: %s\n", project.Description)
		fmt.Fprintf(out, "Created:     %s\n", project.CreatedAt.Format("2006-01-02 15:04:05"))
		return nil
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd, projectsGetCmd)
	rootCmd.AddCommand(projectsCmd)
}
