package components

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	prevHash     [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		prevHash:     prevHash,
		timestamp:    time.Now().UnixNano(),
		transactions: transactions,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("prevHash: %x\n", b.prevHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := b.MarshalJSON()
	// fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PrevHash     [32]byte       `json:"prev_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}
