package main

import (
  "context"
  "fmt"
  "log"
  "math/big"

  "github.com/ethereum/go-ethereum/ethclient"
)
func main() {
  client, err := ethclient.Dial("https://mainnet.infura.io")
  if err != nil {
    log.Fatal(err)
  }

  header, err := client.HeaderByNumber(context.Background(), nil)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(header.Number.String())

  blockNumber := big.NewInt(5671744)
  block, err := client.BlockByNumber(context.Background(), blockNumber)
  if err != nil {
    log.Fatal(err)
  }

  // block's information
  fmt.Println(block.Number().Uint64())
  fmt.Println(block.Time().Uint64())
  fmt.Println(block.Difficulty().Uint64())
  fmt.Println(block.Hash().Hex())
  fmt.Println(len(block.Transactions()))

  // transactions in one block
  count, err := client.TransactionCount(context.Background(), block.Hash())
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(count)
}