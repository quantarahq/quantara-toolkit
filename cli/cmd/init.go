package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type projectManifest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	initName        string
	initDescription string
)

var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Scaffold a new Quantara project directory locally",
	Long: `Creates a local project directory with a quantara.json manifest.
This is purely local scaffolding — it does not talk to quantara-core.
Use "quantara deploy" (not yet implemented) to register it with a
running quantara-core instance.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "."
		if len(args) == 1 {
			dir = args[0]
		}

		name := initName
		if name == "" {
			abs, err := filepath.Abs(dir)
			if err != nil {
				return err
			}
			name = filepath.Base(abs)
		}

		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("creating project directory: %w", err)
		}

		manifestPath := filepath.Join(dir, "quantara.json")
		if _, err := os.Stat(manifestPath); err == nil {
			return fmt.Errorf("%s already exists", manifestPath)
		}

		manifest := projectManifest{Name: name, Description: initDescription}
		data, err := json.MarshalIndent(manifest, "", "  ")
		if err != nil {
			return err
		}
		data = append(data, '\n')

		if err := os.WriteFile(manifestPath, data, 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", manifestPath, err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Scaffolded %s\n", manifestPath)
		return nil
	},
}

func init() {
	initCmd.Flags().StringVar(&initName, "name", "", "project name (defaults to the directory name)")
	initCmd.Flags().StringVar(&initDescription, "description", "", "project description")
	rootCmd.AddCommand(initCmd)
}
