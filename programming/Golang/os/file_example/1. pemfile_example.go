package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {

	var fname string
	timestmp := time.Now().Format("2006-01-02T15:04:05.000Z")

	//-------------------------------------------------
	// 1. GenerateKeyPair (ed25519)
	pub, prv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Printf("error : failed to GenerateKey(%s)\n", err.Error())
		return
	}

	b_prv, err := x509.MarshalPKCS8PrivateKey(prv)
	if err != nil {
		log.Printf("error : failed to MarshalPKCS8PrivateKey(%s)\n", err.Error())
		return
	}

	b_pub, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		log.Printf("error : failed to MarshalPKIXPublicKey(%s)\n", err.Error())
		return
	}

	//-------------------------------------------------
	// 2. Generate PEM Blocks
	pubKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b_pub,
	}

	prvKeyPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b_prv,
	}

	//-------------------------------------------------
	// 3. Create filename and Mkdir (as current os ...)
	fname = fmt.Sprintf("./output/%x", sha256.Sum256([]byte(timestmp)))
	err = os.MkdirAll(filepath.Dir(fname), os.ModePerm)
	if err != nil {
		log.Printf("error : failed to mkdir(%s)\n", err.Error())
		return
	}

	//-------------------------------------------------
	// 4. Create file
	pubKeyFile, err := os.Create(fname + "_pub.pem")
	if err != nil {
		log.Printf("error : failed to os.Create(%s)\n", err.Error())
		return
	}
	defer pubKeyFile.Close()
	err = pem.Encode(pubKeyFile, pubKeyPEM)
	if err != nil {
		log.Printf("error : failed to encode PEM(%s)\n", err.Error())
		return
	}

	prvKeyFile, err := os.Create(fname + "_prv.pem")
	if err != nil {
		log.Printf("error : failed to os.Create(%s)\n", err.Error())
		return
	}
	defer prvKeyFile.Close()
	err = pem.Encode(prvKeyFile, prvKeyPEM)
	if err != nil {
		log.Printf("error : failed to encode PEM(%s)\n", err.Error())
		return
	}

	//-------------------------------------------------
	log.Printf("Success (fname_dir : %s)\n", fname)
	return
}
