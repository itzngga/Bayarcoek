package src

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
	"github.com/gossl/rsam"
	"github.com/itzngga/Bayarcoek/model"
	"io"
	"time"

	"mime/multipart"
	"net/http"
	"strings"
)

func UploadKeyToAnonfiles(pubKey *rsa.PublicKey) (string, bool) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = "Upload key to anonfiles..."
	s.Start()
	reader := bytes.NewReader(rsam.PublicKeyToBytes(pubKey))
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "thekey")
	io.Copy(part, reader)
	writer.Close()

	r, _ := http.NewRequest("POST", "https://api.anonfiles.com/upload", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return "", false
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return "", false
	}

	resp := new(model.AnonfilesResult)
	json.Unmarshal(b, &resp)

	if !resp.Status {
		return "", false
	}

	s.Stop()
	return resp.Data.File.Url.Short, true
}

func AnonfilesToPubkey(url string) (*rsa.PublicKey, bool) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = "Getting Keys..."
	s.Start()
	if !strings.HasPrefix(url, "https://anonfiles.com/") {
		fmt.Printf("ERROR: %s\n%s\n", url, "INVALID KEY URL")
		return nil, false
	}
	r, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil, false
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	key, ok := doc.Find("#download-url").Attr("href")
	if !ok {
		fmt.Printf("ERROR: %s\n%s\n", url, "INVALID KEY URL")
		return nil, false
	}

	req, _ := http.NewRequest("GET", key, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil, false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil, false
	}

	defer resp.Body.Close()

	publicKey, err := rsam.BytesToPublicKey(body)
	if err != nil {
		fmt.Printf("ERROR: %s\n%s\n", url, "INVALID KEY FILE")
		return nil, false
	}
	s.Stop()
	return publicKey, true
}
