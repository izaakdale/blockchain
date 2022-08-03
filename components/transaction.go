package components

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress   string
	recipientBlockchainAddres string
	value                     float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("%s ===> %s @ %f IZC\n", t.senderBlockchainAddress, t.recipientBlockchainAddres, t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender"`
		Recipient string  `json:"recipient"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddres,
		Value:     t.value,
	})
}
