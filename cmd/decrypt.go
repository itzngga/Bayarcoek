package cmd

import (
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Aliases: []string{"dec", "dc", "d"},
	Short:   "Decrypt files or directories.",
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
