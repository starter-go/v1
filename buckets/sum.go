package buckets

import (
	"crypto"
)

type HashAlgorithm = crypto.Hash

const (
	MD4 HashAlgorithm = crypto.MD4
	MD5 HashAlgorithm = crypto.MD5

	SHA1   HashAlgorithm = crypto.SHA1
	SHA224 HashAlgorithm = crypto.SHA224
	SHA256 HashAlgorithm = crypto.SHA256
	SHA384 HashAlgorithm = crypto.SHA384
	SHA512 HashAlgorithm = crypto.SHA512

	MD5SHA1   HashAlgorithm = crypto.MD5SHA1
	RIPEMD160 HashAlgorithm = crypto.RIPEMD160

	SHA3_224   HashAlgorithm = crypto.SHA3_224
	SHA3_256   HashAlgorithm = crypto.SHA3_256
	SHA3_384   HashAlgorithm = crypto.SHA3_384
	SHA3_512   HashAlgorithm = crypto.SHA3_512
	SHA512_224 HashAlgorithm = crypto.SHA512_224
	SHA512_256 HashAlgorithm = crypto.SHA512_256

	BLAKE2s_256 HashAlgorithm = crypto.BLAKE2s_256
	BLAKE2b_256 HashAlgorithm = crypto.BLAKE2b_256
	BLAKE2b_384 HashAlgorithm = crypto.BLAKE2b_384
	BLAKE2b_512 HashAlgorithm = crypto.BLAKE2b_512
)
