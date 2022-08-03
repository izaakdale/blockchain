package components

import (
	"fmt"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
)

func (bc *Blockchain) ValidateProof(nonce int, prevHash [32]byte, txs []*Transaction, difficulty int) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := Block{nonce, prevHash, 0, txs}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeroes
}

func (bc *Blockchain) ProofOfWork() int {
	txs := bc.CopyTransactionPool()
	prevHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidateProof(nonce, prevHash, txs, MINING_DIFFICULTY) {
		nonce++
	}
	return nonce
}
