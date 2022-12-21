package main

import (
	"bytes"
	"fmt"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	//1. bip-39로 seed 만들고
	//2. bip-32로 child key 뽑아내고
	//3. key복구하려면 mnemonic code 를 알고있어야 할듯

	//-------------------------------------------------------------
	//1.bip-39로 seed 생성
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	fmt.Println("1. Mnemonic: ", mnemonic)

	seed := bip39.NewSeed(mnemonic, "password1234") //"Secret Passphrase")
	//-------------------------------------------------------------

	//-------------------------------------------------------------
	//2.bip-32로 child key 뽑아내기
	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	fmt.Println("2. Master private key: ", masterKey)
	fmt.Println("2. Master public key: ", publicKey)

	departmentKeys := map[string]*bip32.Key{}
	departmentKeys["Sales"], _ = masterKey.NewChildKey(0)
	departmentKeys["Marketing"], _ = masterKey.NewChildKey(1)
	departmentKeys["Engineering"], _ = masterKey.NewChildKey(2)
	departmentKeys["Customer Support"], _ = masterKey.NewChildKey(3)

	fmt.Println("2-1. ChildKey[0] prvkey : ", departmentKeys["Sales"])
	fmt.Println("2-1. ChildKey[0] pubkey : ", departmentKeys["Sales"].PublicKey())
	fmt.Println("2-2. ChildKey[1] prvkey : ", departmentKeys["Marketing"])
	fmt.Println("2-2. ChildKey[1] pubkey : ", departmentKeys["Marketing"].PublicKey())
	fmt.Println("2-3. ChildKey[2] prvkey : ", departmentKeys["Engineering"])
	fmt.Println("2-3. ChildKey[2] pubkey : ", departmentKeys["Engineering"].PublicKey())
	fmt.Println("2-4. ChildKey[3] prvkey : ", departmentKeys["Customer Support"])
	fmt.Println("2-4. ChildKey[3] pubkey : ", departmentKeys["Customer Support"].PublicKey())

	//-------------------------------------------------------------
	//-------------------------------------------------------------

	fmt.Println("============Derive from Mnemonic code============")

	fmt.Println("============          (FAIL)         ============")
	//3. mnemonic code로 key만들기
	input_mnemonic := "mnemonic_example" //mnemonic
	input_passwd := "password1234"

	seed_check := bip39.NewSeed(input_mnemonic, input_passwd)

	GenmasterKey, _ := bip32.NewMasterKey(seed_check)
	GenpublicKey := GenmasterKey.PublicKey()

	fmt.Println("3. Generate Master private key: ", GenmasterKey)
	fmt.Println("3. Generate Master public key: ", GenpublicKey)

	departmentKeys2 := map[string]*bip32.Key{}
	departmentKeys2["Sales"], _ = GenmasterKey.NewChildKey(0)
	departmentKeys2["Marketing"], _ = GenmasterKey.NewChildKey(1)
	departmentKeys2["Engineering"], _ = GenmasterKey.NewChildKey(2)
	departmentKeys2["Customer Support"], _ = GenmasterKey.NewChildKey(3)

	fmt.Println("3-1. ChildKey[0] prvkey : ", departmentKeys2["Sales"])
	fmt.Println("3-1. ChildKey[0] pubkey : ", departmentKeys2["Sales"].PublicKey())
	real_s_childkey0, _ := departmentKeys["Sales"].Serialize()
	check_s_childkey0, _ := departmentKeys2["Sales"].Serialize()
	fmt.Println("IsReal ChildKey[0]? : ", bytes.Compare(real_s_childkey0, check_s_childkey0))

	fmt.Println("3-2. ChildKey[1] prvkey : ", departmentKeys2["Marketing"])
	fmt.Println("3-2. ChildKey[1] pubkey : ", departmentKeys2["Marketing"].PublicKey())
	real_s_childkey1, _ := departmentKeys["Marketing"].Serialize()
	check_s_childkey1, _ := departmentKeys2["Marketing"].Serialize()
	fmt.Println("IsReal ChildKey[1]? : ", bytes.Compare(real_s_childkey1, check_s_childkey1))

	fmt.Println("3-3. ChildKey[2] prvkey : ", departmentKeys2["Engineering"])
	fmt.Println("3-3. ChildKey[2] pubkey : ", departmentKeys2["Engineering"].PublicKey())
	real_s_childkey2, _ := departmentKeys["Engineering"].Serialize()
	check_s_childkey2, _ := departmentKeys2["Engineering"].Serialize()
	fmt.Println("IsReal ChildKey[2]? : ", bytes.Compare(real_s_childkey2, check_s_childkey2))

	fmt.Println("3-4. ChildKey[3] prvkey : ", departmentKeys2["Customer Support"])
	fmt.Println("3-4. ChildKey[3] pubkey : ", departmentKeys2["Customer Support"].PublicKey())
	real_s_childkey3, _ := departmentKeys["Customer Support"].Serialize()
	check_s_childkey3, _ := departmentKeys2["Customer Support"].Serialize()
	fmt.Println("IsReal ChildKey[3]? : ", bytes.Compare(real_s_childkey3, check_s_childkey3))

	//-------------------------------------------------------------
	//-------------------------------------------------------------

	fmt.Println("============        (SUCCESS)        ============")
	input_mnemonic = mnemonic
	input_passwd = "password1234"

	seed_check = bip39.NewSeed(input_mnemonic, input_passwd)

	GenmasterKey, _ = bip32.NewMasterKey(seed_check)
	GenpublicKey = GenmasterKey.PublicKey()

	fmt.Println("4. Generate Master private key: ", GenmasterKey)
	fmt.Println("4. Generate Master public key: ", GenpublicKey)

	departmentKeys2 = map[string]*bip32.Key{}
	departmentKeys2["Sales"], _ = GenmasterKey.NewChildKey(0)
	departmentKeys2["Marketing"], _ = GenmasterKey.NewChildKey(1)
	departmentKeys2["Engineering"], _ = GenmasterKey.NewChildKey(2)
	departmentKeys2["Customer Support"], _ = GenmasterKey.NewChildKey(3)

	fmt.Println("4-1. ChildKey[0] prvkey : ", departmentKeys2["Sales"])
	fmt.Println("4-1. ChildKey[0] pubkey : ", departmentKeys2["Sales"].PublicKey())
	real_s_childkey0, _ = departmentKeys["Sales"].Serialize()
	check_s_childkey0, _ = departmentKeys2["Sales"].Serialize()
	fmt.Println("IsReal ChildKey[0]? : ", bytes.Compare(real_s_childkey0, check_s_childkey0))

	fmt.Println("4-2. ChildKey[1] prvkey : ", departmentKeys2["Marketing"])
	fmt.Println("4-2. ChildKey[1] pubkey : ", departmentKeys2["Marketing"].PublicKey())
	real_s_childkey1, _ = departmentKeys["Marketing"].Serialize()
	check_s_childkey1, _ = departmentKeys2["Marketing"].Serialize()
	fmt.Println("IsReal ChildKey[1]? : ", bytes.Compare(real_s_childkey1, check_s_childkey1))

	fmt.Println("4-3. ChildKey[2] prvkey : ", departmentKeys2["Engineering"])
	fmt.Println("4-3. ChildKey[2] pubkey : ", departmentKeys2["Engineering"].PublicKey())
	real_s_childkey2, _ = departmentKeys["Engineering"].Serialize()
	check_s_childkey2, _ = departmentKeys2["Engineering"].Serialize()
	fmt.Println("IsReal ChildKey[2]? : ", bytes.Compare(real_s_childkey2, check_s_childkey2))

	fmt.Println("4-4. ChildKey[3] prvkey : ", departmentKeys2["Customer Support"])
	fmt.Println("4-4. ChildKey[3] pubkey : ", departmentKeys2["Customer Support"].PublicKey())
	real_s_childkey3, _ = departmentKeys["Customer Support"].Serialize()
	check_s_childkey3, _ = departmentKeys2["Customer Support"].Serialize()
	fmt.Println("IsReal ChildKey[3]? : ", bytes.Compare(real_s_childkey3, check_s_childkey3))

	//compare output
	// a==b : 0
	// a<b : -1
	// b>a : 1
	/*

		type Key struct {
			Key         []byte // 33 bytes
			Version     []byte // 4 bytes
			ChildNumber []byte // 4 bytes
			FingerPrint []byte // 4 bytes
			ChainCode   []byte // 32 bytes
			Depth       byte   // 1 bytes
			IsPrivate   bool   // unserialized
		}

	*/
}
