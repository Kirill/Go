package test1

import "crypto/sha512"

func f(s []byte, i int) []byte {
	if i == 0 {
		v := sha512.Sum512(s)
		return v[:]
	}
	return g(s, i)
}

func g(s []byte, i int) []byte {
	switch i {
	case 512:
		v := sha512.Sum512(s)
		return v[:]
	case 384:
		v := sha512.Sum384(s)
		return v[:]
	case 256:
		v := sha512.Sum512_256(s)
		return v[:]
	case 224:
		v := sha512.Sum512_224(s)
		return v[:]
	default:
		panic("unknown base")
	}
}
