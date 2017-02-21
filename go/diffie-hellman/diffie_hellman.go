package diffiehellman

import "math/big"

const testVersion = 1

// PrivateKey generates a private key from a prime number
func PrivateKey(p *big.Int) *big.Int {
	return &big.Int{}
}

// PublicKey generates a public key from a private key and two prime numbers.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return &big.Int{}
}

// NewPair generates a new pair of keys
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	return &big.Int{}, &big.Int{}
}

// SecretKey generates a secret key
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return &big.Int{}
}
