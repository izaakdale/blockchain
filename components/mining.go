package components

import "log"

const (
	MINING_SENDER = "IZCHAIN"
	MINING_REWARD = 1.0
)

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockChainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	prevHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, prevHash)
	log.Println("action=mining, status=success!")
	return true
}
