package main

// A keystore is a file containing an encrypted wallet private key.
// Keystores in go-ethereum can only contain one wallet key pair per file.
import (
  "fmt"
  "io/ioutil"
  "log"
  "os"

  "github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
  ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
  password := "secret"
  account, err := ks.NewAccount(password)
  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Println(account.Address.Hex())
}

func importKs() {
  file := "./wallets/UTC--2019-02-03T15-35-42.555415000Z--a9212c5333cfd32c201d99c360fdf1f713b85108"
  ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
  jsonBytes, err := ioutil.ReadFile(file)
  if err != nil {
    log.Fatal(err)
  }
  password := "secret"
  account, err := ks.Import(jsonBytes, password, password)
  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Println(account.Address.Hex())
  if err := os.Remove(file); err != nil {
    log.Fatal(err)
  }
}
func main() {
  // createKs()
  importKs()
}