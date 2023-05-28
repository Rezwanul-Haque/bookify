package cmd

import (
	"bookify/infra/config"
	"bookify/infra/conn/db"
	"bookify/infra/logger"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "bookify",
		Short: "implementing book library",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	config.LoadConfig()
	logger.NewLogClient(config.App().LogLevel)
	lc := logger.Client()
	db.NewDbClient(lc)

	lc.Info("about to start the application")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
