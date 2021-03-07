package cmd

import (
	"fmt"
	"os"

	"github.com/jsws/certutil/pkg/connect"
	"github.com/jsws/certutil/pkg/file"
	"github.com/jsws/certutil/pkg/info"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about a certificate chain.",
}

func infoConnect(cmd *cobra.Command, args []string) {
	host := args[0]
	certs, err := connect.Connect(host, serverName, port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if serverName == "" {
		serverName = host
	}
	info.PrintInfo(certs, serverName)
}

func infoRead(cmd *cobra.Command, args []string) {
	filepath := args[0]
	certs, err := file.ReadPEM(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	info.PrintInfo(certs, serverName)
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(newConnectCmd(infoConnect))
	infoCmd.AddCommand(newReadCmd(infoRead))
}
