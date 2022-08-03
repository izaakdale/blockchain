package main

import (
	"log"

	cmp "github.com/izaakdale/blockchain/components"
)

func init() {
	log.SetPrefix("Blockchain : ")
}

func main() {
	bc := cmp.NewBlockchain()
	b := bc.LastBlock()
	bc.CreateBlock(5, b.Hash())
	bc.AddTransaction("izaak", "mahtab", 0.129804)
	b = bc.LastBlock()
	bc.CreateBlock(10, b.Hash())
	bc.AddTransaction("mahtab", "izaak", 0.30080)
	bc.AddTransaction("mahtab", "izaak", 0.50080)
	b = bc.LastBlock()
	bc.CreateBlock(15, b.Hash())

	bc.Print()
	// fmt.Println(bc.chain[0].Hash())
}
