package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// subscribe to get events when their is a new block mined.
func main() {
	// using websocket
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	// create new channel
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// the subscription will push new block headers to our channel
	// so we'll use a select statement to listen for new messages.
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Hash: ", block.Hash().Hex())
			fmt.Println("Number: ", block.Number().Uint64())
			fmt.Println("Time: ", block.Time().Uint64())
			fmt.Println("Nonce: ", block.Nonce())
			fmt.Println("Transaction len: ", len(block.Transactions()))
			fmt.Println("-------------------------------------")
		}
	}
}
