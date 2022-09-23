package cmd

import (
	"crypto/rsa"
	"fmt"
	"github.com/gossl/rsam"
	"github.com/itzngga/Bayarcoek/src"
	"github.com/spf13/cobra"
)

// encryptFilesCmd represents the directory command
var encryptFilesCmd = &cobra.Command{
	Use:     "file",
	Aliases: []string{"files", "f", "fs"},
	Short:   "Encrypt one or more files.",
	Long:    `Encrypt one or more files, Default keys will be generated automatically if string keys are not provided.`,
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
		if privateKey == "" && privateKeyPath != "" {
			priv, ok := src.GetPrivateKeyFromPath(privateKeyPath)
			if !ok {
				return
			}
			pvKey = priv
		}
		if publicKey == "" && publicKeyPath != "" {
			pub, ok := src.GetPublicKeyFromPath(publicKeyPath)
			if !ok {
				return
			}
			pubKey = pub
		}
		if privateKey == "" && privateKeyPath == "" {
			fmt.Println("ERROR: Missing private key.")
			return
		}
		if generateUrl && publicKey == "" && publicKeyPath == "" {
			fmt.Println("ERROR: Missing public key.")
			return
		}
		for _, arg := range args {
			err := src.MakeSingleBayarCoek(arg, extension, overwrite, pvKey)
			if err != nil {
				return
			}
		}
		if generateUrl {
			result, ok := src.UploadKeyToAnonfiles(pubKey)
			if ok {
				fmt.Println("SUCESS:\nUploaded Key: ", result)
			}
		}
	},
}

func init() {
	encryptCmd.AddCommand(encryptFilesCmd)

	encryptFilesCmd.Flags().StringVar(&privateKeyPath, "privPath", "keys/private.key", "Path to private key file.")
	encryptFilesCmd.Flags().StringVar(&publicKeyPath, "pubPath", "keys/public.key", "Path to public key file.")
	encryptFilesCmd.Flags().StringVar(&privateKey, "privateKey", "", "Private key string.")
	encryptFilesCmd.Flags().StringVar(&publicKey, "publicKey", "", "Public key string.")
	encryptFilesCmd.Flags().StringVarP(&extension, "extension", "e", "bayarcoek", "Encrypted file custom extension.")
	encryptFilesCmd.Flags().BoolVarP(&generateUrl, "url", "u", false, "Generate key url.")
	encryptFilesCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", true, "Overwrite existing files.")
}
