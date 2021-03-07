package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// newReadCmd creates a new (sub)command for reading a PEM file. Executing the suplied f.
func newReadCmd(f func(cmd *cobra.Command, args []string)) *cobra.Command {
	readCmd := &cobra.Command{
		Use:   "read filepath.pem",
		Short: "Reads from a file.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("Requires filepath e.g. certs/file.pem")
			}
			return nil
		},
		Run: f,
	}

	return readCmd
}
