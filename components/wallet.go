package components

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockChainAddress string
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	w.AssignAddress()

	return w
}

func (w *Wallet) AssignAddress() {

	// perform sha256 hashing on public key
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// perform ripemd 160 hashing on the result
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// prepend version byte on ripemd result
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// sha256 hash of merged verion and ripemd byte slices
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	// sha256 hash again
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	// take first for bytes for checksum
	chsum := digest6[:4]
	// take bytes and append to merged ripemd bytes
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	// convert to byte string in base58
	address := base58.Encode(dc8)

	w.blockChainAddress = string(address)
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}
func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X, w.publicKey.Y)
}
func (w *Wallet) Address() string {
	return w.blockChainAddress
}
