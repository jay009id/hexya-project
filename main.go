// This file is autogenerated by hexya-server
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package main

import (
	"fmt"
	"os"

	_ "github.com/hexya-addons/web"
	"github.com/hexya-erp/hexya/cmd"
	_ "github.com/jay009id/hexya-project/openacademy"
	"github.com/spf13/cobra"
)

func main() {
	var hexyaCmd = &cobra.Command{
		Use:   ".",
		Short: "Hexya is an open source modular ERP",
		Long:  "Hexya is an open source modular ERP written in Go. It is designed for high demand business data processing while being easily customizable",
	}
	cmd.SetHexyaFlags(hexyaCmd)

	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the Hexya server",
		Long:  "Start the Hexya server",
		Run: func(c *cobra.Command, args []string) {
			cmd.StartServer()
		},
	}
	hexyaCmd.AddCommand(serverCmd)
	cmd.SetServerFlags(serverCmd)

	var updateDBCmd = &cobra.Command{
		Use:   "updatedb",
		Short: "Update the database schema",
		Long:  "Synchronize the database schema with the models definitions.",
		Run: func(c *cobra.Command, args []string) {
			cmd.UpdateDB()
		},
	}
	hexyaCmd.AddCommand(updateDBCmd)

	cobra.OnInitialize(cmd.InitConfig)

	if err := hexyaCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
