package src

import (
	"crypto/rsa"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/gossl/rsam"
	"os"
	"time"
)

func GenerateBayarcoekKeys(privPath, pubPath string) (*rsa.PrivateKey, *rsa.PublicKey) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = "Generating Keys..."
	s.Start()

	var priv *rsa.PrivateKey
	var pub *rsa.PublicKey

	_, err := os.Stat(privPath)
	if os.IsNotExist(err) {
		s.Suffix = "Generating New a Pair Private Key and Public Key..."
		s.Restart()
		priv, pub, err = rsam.GeneratePairKeys(2048)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		_ = os.Mkdir("keys", 0777)

		err = os.WriteFile(privPath, rsam.PrivateKeyToBytes(priv), 0777)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		err = os.WriteFile(pubPath, rsam.PublicKeyToBytes(pub), 0777)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	} else {
		s.Suffix = "Reading Keys From Given Paths..."
		s.Restart()
		privFile, err := os.ReadFile(privPath)
		priv, err = rsam.BytesToPrivateKey(privFile)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		pubFile, err := os.ReadFile(pubPath)
		pub, err = rsam.BytesToPublicKey(pubFile)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}

	s.Stop()
	return priv, pub
}

func GetPublicKeyFromPath(path string) (*rsa.PublicKey, bool) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = "Getting Public Key..."
	s.Start()

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("ERROR: %s\n%s\n", path, "PUBLIC KEY NOT FOUND.")
		return nil, false
	}

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil, false
	}

	pub, err := rsam.BytesToPublicKey(file)
	if err != nil {
		fmt.Printf("ERROR: %s\n%s\n", path, "INVALID PUBLIC KEY.")
		return nil, false
	}
	s.Stop()
	return pub, true
}

func GetPrivateKeyFromPath(path string) (*rsa.PrivateKey, bool) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = "Getting Private Key..."
	s.Start()

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("ERROR: %s\n%s\n", path, "PRIVATE KEY NOT FOUND.")
		return nil, false
	}

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil, false
	}

	priv, err := rsam.BytesToPrivateKey(file)
	if err != nil {
		fmt.Printf("ERROR: %s\n%s\n", path, "INVALID PRIVATE KEY.")
		return nil, false
	}
	s.Stop()
	return priv, true
}
