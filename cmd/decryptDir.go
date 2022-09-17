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

// decryptDirCmd represents the decryptDir command
var decryptDirCmd = &cobra.Command{
	Use:     "directory",
	Aliases: []string{"dir", "d", "dr"},
	Short:   "Decrypt one or more directories.",
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
			err := src.DecryptDirBayarcoek(arg, extension, pubKey)
			if err != nil {
				return
			}
		}
	},
}

func init() {
	decryptCmd.AddCommand(decryptDirCmd)

	decryptDirCmd.Flags().StringVar(&publicKey, "publicKey", "", "Public key string.")
	decryptDirCmd.Flags().StringVarP(&url, "url", "u", "", "Input anonfiles key url.")
	decryptDirCmd.Flags().StringVar(&publicKeyPath, "pubPath", "keys/public.key", "Path to public key file.")
	decryptDirCmd.Flags().StringVarP(&extension, "extension", "e", "bayarcoek", "Encrypted file custom extension.")
}
