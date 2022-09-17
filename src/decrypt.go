package src

import (
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gossl/rsam"
	"os"
	"path/filepath"
	"strings"
)

func DecryptDirBayarcoek(dirPath, extension string, publicKey *rsa.PublicKey) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if !strings.HasSuffix(path, "."+extension) {
				return errors.New(path + "\nINVALID ENCRYPTED FILE")
			}
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			ciphertext, err := rsam.DecryptWithPublicKey(file, publicKey, sha256.New())
			if err != nil {
				return errors.New(path + "\nINVALID ENCRYPTED FILE")
			}
			originName := strings.Split(string(ciphertext), "🤣🤣🤣")
			if len(originName) == 1 {
				return errors.New(path + "\nINVALID ENCRYPTED FILE")
			}
			err = os.WriteFile(path, []byte(originName[0]), 0777)
			err = os.Rename(path, originName[1])
			if err != nil {
				return err
			}
		} else {
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	fmt.Println(fmt.Sprintf("SUCCESS: %s DECRYPTED", dirPath))
	return nil
}

func DecryptSingleBayarcoek(filePath, extension string, publicKey *rsa.PublicKey) error {
	info, err := os.Lstat(filePath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	if info.IsDir() {
		fmt.Printf("ERROR: %s\n%s\n", filePath, "FILEPATH IS DIRECTORY")
		return errors.New(fmt.Sprintf("ERROR: %s\n%s\n", filePath, "FILEPATH IS DIRECTORY"))
	}
	if !strings.HasSuffix(filePath, "."+extension) {
		fmt.Printf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE")
		return errors.New(fmt.Sprintf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE"))
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	ciphertext, err := rsam.DecryptWithPublicKey(file, publicKey, sha256.New())
	if err != nil {
		fmt.Printf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE")
		return errors.New(fmt.Sprintf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE"))
	}
	originName := strings.Split(string(ciphertext), "🤣🤣🤣")
	if len(originName) == 1 {
		fmt.Printf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE")
		return errors.New(fmt.Sprintf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE"))
	}
	err = os.WriteFile(filePath, []byte(originName[0]), 0777)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	err = os.Rename(filePath, originName[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	fmt.Println(fmt.Sprintf("SUCCESS: %s DECRYPTED", filePath))
	return nil
}
