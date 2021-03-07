package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	host       string
	serverName string
	port       uint16
)

// newConnectCmd creates a new (sub)command for connecting to a server. Executing the suplied f.
func newConnectCmd(f func(cmd *cobra.Command, args []string)) *cobra.Command {
	connectCmd := &cobra.Command{
		Use:   "connect host",
		Short: "Connects to a server.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("Requires a single host e.g. badssl.com")
			}
			// TODO: Should accept just a https URL here for convenience too.
			return nil

		},
		Run: f,
	}
	connectCmd.Flags().StringVarP(&serverName, "servername", "s", "", "SNI for virtual servers (if not set host will be used)")
	connectCmd.Flags().Uint16VarP(&port, "port", "p", 443, "Port to connect to")

	return connectCmd
}
