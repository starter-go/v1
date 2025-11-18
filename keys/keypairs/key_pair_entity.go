package keypairs

import "github.com/starter-go/v1/keys"

type PublicKey interface {
	keys.Entity
}

type PrivateKey interface {
	keys.Entity
}

type Pair interface {
	keys.Entity

	GetPublicKey() PublicKey

	GetPrivateKey() PrivateKey
}
