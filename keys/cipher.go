package keys

type CipherContext struct {
	Padding   Padding
	Mode      BlockMode
	IV        []byte
	Plain     []byte
	Encrypted []byte
	Key       KeyEntity
}

type Cipher interface {
	Ecrypt(cc *CipherContext) error

	Derypt(cc *CipherContext) error
}
