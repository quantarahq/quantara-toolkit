// Command quantara is the Quantara CLI — an alternative client to
// quantara-core's REST API, the same one quantara-web uses.
package main

import (
	"fmt"
	"os"

	"github.com/quantarahq/quantara-toolkit/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
