package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	privateKeyPath string
	publicKeyPath  string
	privateKey     string
	publicKey      string
	extension      string
	verbose        bool
	generateUrl    bool
	url            string
	overwrite      bool
)

var rootCmd = &cobra.Command{
	Use:   "bayarcoek",
	Short: "Encrypt your project to secure transactions.",
	Long:  `Use this command to encrypt your project to secure transactions between client and buyer.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
