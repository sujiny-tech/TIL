package main

import (
	"fmt"
	"log"

	"github.com/tyler-smith/go-bip32"
)

// Example address creation for a fictitious company ComputerVoice Inc. where
// each department has their own wallet to manage
func test1() {
	// Generate a seed to determine all keys from.
	// This should be persisted, backed up, and secured
	seed, err := bip32.NewSeed()
	if err != nil {
		log.Fatalln("Error generating seed:", err)
	}

	// Create master private key from seed
	computerVoiceMasterKey, _ := bip32.NewMasterKey(seed)

	// Map departments to keys
	// There is a very small chance a given child index is invalid
	// If so your real program should handle this by skipping the index
	departmentKeys := map[string]*bip32.Key{}
	departmentKeys["Sales"], _ = computerVoiceMasterKey.NewChildKey(0)
	departmentKeys["Marketing"], _ = computerVoiceMasterKey.NewChildKey(1)
	departmentKeys["Engineering"], _ = computerVoiceMasterKey.NewChildKey(2)
	departmentKeys["Customer Support"], _ = computerVoiceMasterKey.NewChildKey(3)

	// Create public keys for record keeping, auditors, payroll, etc
	departmentAuditKeys := map[string]*bip32.Key{}
	departmentAuditKeys["Sales"] = departmentKeys["Sales"].PublicKey()
	departmentAuditKeys["Marketing"] = departmentKeys["Marketing"].PublicKey()
	departmentAuditKeys["Engineering"] = departmentKeys["Engineering"].PublicKey()
	departmentAuditKeys["Customer Support"] = departmentKeys["Customer Support"].PublicKey()

	// Print public keys
	for department, pubKey := range departmentAuditKeys {
		fmt.Printf("[%s] : ", department)
		fmt.Println(pubKey)
	}

	// Print public keys
	for department, prvKey := range departmentKeys {
		fmt.Printf("[%s] : ", department)
		fmt.Println(prvKey)
	}
}
