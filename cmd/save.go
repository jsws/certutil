package cmd

import (
	"fmt"
	"os"

	"github.com/jsws/certutil/pkg/connect"
	"github.com/jsws/certutil/pkg/file"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves a certificate presented by a TLS server as a PEM encoded file.",
}

func saveConnect(cmd *cobra.Command, args []string) {
	host := args[0]
	certs, err := connect.Connect(host, serverName, port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer, err := file.GetWriter(outputFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = file.WritePEM(certs, writer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	outputFilePath string
)

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.AddCommand(newConnectCmd(saveConnect))
	saveCmd.PersistentFlags().StringVarP(&outputFilePath, "output", "o", "", "File path to save file")
}
