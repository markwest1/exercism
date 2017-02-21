package diffiehellman

import (
	"crypto/rand"
	"log"
	"math/big"
)

const testVersion = 1

var uno = big.NewInt(1)

// PrivateKey generates a private key in the range (1, p)
func PrivateKey(p *big.Int) *big.Int {
	private, err := rand.Int(rand.Reader, p)
	if err != nil {
		log.Fatalf("PrivateKey error: %s", err.Error())
	}

	// Lower limit is 1 (non-inclusive)
	for private.Cmp(uno) <= 0 {
		private.Add(private, uno)
	}

	return private
}

// PublicKey generates a public key from a private key:
// public-key = g**private-key mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	public := &big.Int{}
	return public.Exp(big.NewInt(g), private, p)
}

// NewPair generates a new pair (private+public) of keys
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey generates a secret key:
// secret-key = public2**private1 mod p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	secret := &big.Int{}
	return secret.Exp(public2, private1, p)
}
