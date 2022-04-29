package main

import (
	"bytes"
	"context"
	"log"

	client "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	c, err := client.New("https://rpc.cosmos.network:443", "/websocket")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	var (
		vh          []byte
		transitions int
		heights     int64
		height      int64
	)
	height = 5200791
	end := height + 10000

	for height < end {
		res, err := c.BlockchainInfo(ctx, height, 0)
		if err != nil {
			log.Printf("error: %s", err)
			break
		}
		for _, m := range res.BlockMetas {
			heights++
			height++
			if height%100 == 0 {
				log.Printf("height: %d", height)
			}
			h := m.Header.ValidatorsHash
			if !bytes.Equal(h, vh) {
				vh = h
				transitions++
			}
		}
	}
	log.Print("Counting complete!")
	log.Printf("transitions: %d", transitions)
	log.Printf("heights counted: %d", heights)
}
