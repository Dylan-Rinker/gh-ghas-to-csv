package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {

	// Instantiate struct to contain values from cobra flags; arguments are handled within RunE
	// cmdFlags := cmdFlags{}

	// Instantiate cobra command driving work from package
	// Closures are used for cobra command lifecycle hooks for access to cobra flags struct
	cmd := cobra.Command{}

	return &cmd
}
