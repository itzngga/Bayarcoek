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

// encryptDirCmd represents the directory command
var encryptDirCmd = &cobra.Command{
	Use:     "directory",
	Aliases: []string{"dir", "d", "dr"},
	Short:   "Encrypt one or more directory.",
	Long:    `Encrypt one or more directory, Default keys will be generated automatically if string keys are not provided.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			pvKey  *rsa.PrivateKey
			pubKey *rsa.PublicKey
			err    error
		)

		if privateKey != "" {
			pvKey, err = rsam.BytesToPrivateKey([]byte(privateKey))
			if err != nil {
				fmt.Println("ERROR: Invalid given private key.")
				return
			}
		}
		if publicKey != "" {
			pubKey, err = rsam.BytesToPublicKey([]byte(publicKey))
			if err != nil {
				fmt.Println("ERROR: Invalid given public key.")
				return
			}
		}
		if privateKey != "" {
			priv, ok := src.GetPrivateKeyFromPath(privateKey)
			if !ok {
				return
			}
			pvKey = priv
		}
		if publicKey != "" {
			pub, ok := src.GetPublicKeyFromPath(publicKey)
			if !ok {
				return
			}
			pubKey = pub
		}
		for _, arg := range args {
			err := src.MakeDirBayarCoek(arg, extension, pvKey)
			if err != nil {
				return
			}
		}
		if generateUrl {
			result, ok := src.UploadKeyToAnonfiles(pubKey)
			if ok {
				fmt.Println("Uploaded Key: ", result)
			}
		}
	},
}

func init() {
	encryptCmd.AddCommand(encryptDirCmd)

	encryptDirCmd.Flags().StringVar(&privateKeyPath, "privPath", "keys/private.key", "Path to private key file.")
	encryptDirCmd.Flags().StringVar(&publicKeyPath, "pubPath", "keys/public.key", "Path to public key file.")
	encryptDirCmd.Flags().StringVar(&privateKey, "privateKey", "", "Private key string.")
	encryptDirCmd.Flags().StringVar(&publicKey, "publicKey", "", "Public key string.")
	encryptDirCmd.Flags().StringVarP(&extension, "extension", "e", "bayarcoek", "Encrypted file custom extension.")
	encryptDirCmd.Flags().BoolVarP(&generateUrl, "url", "u", false, "Generate key url.")
}
