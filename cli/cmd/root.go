// Package cmd holds the Quantara CLI's Cobra command tree.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var apiURL string

var rootCmd = &cobra.Command{
	Use:   "quantara",
	Short: "Quantara CLI — an alternative client for quantara-core's REST API",
	Long: `Quantara CLI talks to a running quantara-core instance over the same
REST API quantara-web uses — it is an alternative client, not a separate
backend. See https://github.com/quantarahq/quantara-toolkit for the roadmap.`,
	// Runtime errors (API unreachable, 404, etc.) aren't usage mistakes —
	// don't dump the usage block for them, and let main() print the error
	// once instead of Cobra printing it a second time.
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command; main() exits non-zero on error.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	defaultAPIURL := os.Getenv("QUANTARA_API_URL")
	if defaultAPIURL == "" {
		defaultAPIURL = "http://localhost:8080"
	}
	rootCmd.PersistentFlags().StringVar(&apiURL, "api-url", defaultAPIURL,
		"quantara-core API base URL (env: QUANTARA_API_URL)")
}
