/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rsa"
	"fmt"
	"github.com/gossl/rsam"
	"github.com/itzngga/Bayarcoek/src"
	"github.com/spf13/cobra"
)

// decryptFilesCmd represents the decryptFiles command
var decryptFilesCmd = &cobra.Command{
	Use:     "file",
	Aliases: []string{"files", "f", "fs"},
	Short:   "Decrypt one or more files.",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			pubKey *rsa.PublicKey
			err    error
		)

		if url != "" {
			pkey, ok := src.AnonfilesToPubkey(url)
			if !ok {
				return
			}
			pubKey = pkey
		} else if url == "" && publicKeyPath != "" {
			pkey, ok := src.GetPublicKeyFromPath(publicKeyPath)
			if !ok {
				return
			}
			pubKey = pkey
		} else if url == "" && publicKey != "" {
			pubKey, err = rsam.BytesToPublicKey([]byte(publicKey))
			if err != nil {
				fmt.Println("ERROR: Invalid given public key.")
				return
			}
		} else {
			fmt.Println("ERROR: Invalid missing url/public key.")
			return
		}

		for _, arg := range args {
			err := src.DecryptSingleBayarcoek(arg, extension, pubKey)
			if err != nil {
				return
			}
		}
	},
}

func init() {
	decryptCmd.AddCommand(decryptFilesCmd)

	decryptFilesCmd.Flags().StringVar(&publicKey, "publicKey", "", "Public key string.")
	decryptFilesCmd.Flags().StringVarP(&url, "url", "u", "", "Input anonfiles key url.")
	decryptFilesCmd.Flags().StringVar(&publicKeyPath, "pubPath", "keys/public.key", "Path to public key file.")
	decryptFilesCmd.Flags().StringVarP(&extension, "extension", "e", "bayarcoek", "Encrypted file custom extension.")
}
