package keypairs

import "github.com/starter-go/v1/keys"

type SignatureContext struct {
	Algorithm  keys.Algorithm
	PublicKey  PublicKey
	PrivateKey PrivateKey
	Data       []byte
	Signature  []byte
}

type Signer interface {
	Sign(ctx *SignatureContext) error
}

type Verifier interface {
	Verify(ctx *SignatureContext) error
}
