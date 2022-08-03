package main

import (
	"fmt"
	"log"

	cmp "github.com/izaakdale/blockchain/components"
)

func init() {
	log.SetPrefix("Blockchain : ")
}

func main() {
	miner := "miningNerd"
	bc := cmp.NewBlockchain(miner)
	bc.AddTransaction("izaak", "mahtab", 1)
	bc.Mining()

	bc.AddTransaction("mahtab", "izaak", 2)
	bc.AddTransaction("mahtab", "izaak", 3)
	bc.Mining()

	bc.Print()

	fmt.Println(bc.CalcTotalAmount("miningNerd"))
}
