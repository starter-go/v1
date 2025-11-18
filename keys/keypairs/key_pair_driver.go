package keypairs

import "github.com/starter-go/v1/keys"

type Driver interface {
	keys.KeyDriver

	GetSigner() Signer

	GetVerifier() Verifier
}
