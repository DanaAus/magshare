package cmd

import (
	"fmt"
	"os"

	"qshare/internal/handlers"
	"qshare/internal/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qshare",
	Short: "QShare is an instant local network file sharing tool",
	Long:  `QShare allows you to instantly share and receive files across your local network. It spawns ephemeral web servers and provides QR codes for easy access.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If no subcommand is provided, run interactive mode
		cfg, err := ui.RunInteractivePrompts()
		if err != nil {
			return err
		}

		switch cfg.Action {
		case "send":
			opts := handlers.SendOptions{
				Port:   cfg.Port,
				Secure: cfg.Secure,
				PIN:    cfg.PIN,
			}
			return handlers.StartSendServer(cfg.Path, opts)
		case "receive":
			opts := handlers.ReceiveOptions{
				Port:   cfg.Port,
				Secure: cfg.Secure,
				PIN:    cfg.PIN,
			}
			return handlers.StartReceiveServer(cfg.Path, opts)
		case "sync":
			fmt.Println("Sync mode not yet fully implemented in this prototype.")
			return nil
		default:
			return cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can go here
}
