package components

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockChainAddress string
}

func NewBlockchain(blockChainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockChainAddress = blockChainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 25))
	for i, block := range bc.chain {
		fmt.Printf("%s Chain #%d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)

	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddres, t.value),
		)
	}
	return transactions
}

func (bc *Blockchain) CalcTotalAmount(address string) float32 {
	total := float32(0.0)
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			if address == t.recipientBlockchainAddres {
				total += t.value
			}
			if address == t.senderBlockchainAddress {
				total -= t.value
			}
		}
	}
	return total
}
