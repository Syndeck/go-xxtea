// Package xxtea implements XXTEA encryption as defined in Needham and Wheeler's
// 1998 technical report, "Correction to XTEA."
package xxtea

// For details, see http://www.movable-type.co.uk/scripts/xxtea.pdf

import (
	"crypto/cipher"
	"strconv"
)

// The XXTEA block size in bytes.
const BlockSize = 8

// An xxteaCipher is an instance of an XXTEA cipher using a particular key.
type xxteaCipher struct {
	k [4]int32
}

// KeySizeError may be returned by NewCipher.
type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/xtea: invalid key size " + strconv.Itoa(int(k))
}

// NewCipher creates and returns a new cipher.Block. The key argument
// should be the XXTEA key. XXTEA only supports 128 bit (16 byte) keys
// which are converted internally into 4 little-endian uint32 values.
func NewCipher(key []byte) (cipher.Block, error) {
	k := len(key)
	switch k {
	default:
		return nil, KeySizeError(k)
	case 16:
		break
	}
	u := bytesToInt32(key)
	c := new(xxteaCipher)
	copy(c.k[:], u)
	return c, nil
}

func (c *xxteaCipher) BlockSize() int { return BlockSize }

func (c *xxteaCipher) Encrypt(dst, src []byte) {
	v := bytesToInt32(src)
	c.blockEncrypt(v)
	copy(dst, int32ToBytes(v))
}

func (c *xxteaCipher) Decrypt(dst, src []byte) {
	v := bytesToInt32(src)
	c.blockDecrypt(v)
	copy(dst, int32ToBytes(v))
}

const delta = -1640531527

// blockEncrypt encrypts the []uint32 represtentation of a block,
// in-place.
func (c *xxteaCipher) blockEncrypt(v []int32) {
	n := len(v)
	y := v[0]
	z := v[n-1]
	rounds := 6 + 52/n

	var sum int32
	for rounds > 0 {
		rounds--
		sum += delta
		e := (sum >> 2) & 3
		var p int
		for p = 0; p < n-1; p++ {
			y = v[p+1]

			t_a := int32(uint32(z) >> 5) ^ int32(uint32(y) << 2)
			t_b := int32(uint32(y) >> 3) ^ int32(uint32(z) << 4)
			t_c := int32(sum ^ y)
			t_d := (c.k[int32(p & 3) ^ e]) ^ int32(z)

			v[p] += (t_a + t_b) ^ (t_c + t_d)
			z = v[p]
		}
		y = v[0]

		t_a := int32(uint32(z) >> 5) ^ int32(uint32(y) << 2)
		t_b := int32(uint32(y) >> 3) ^ int32(uint32(z) << 4)
		t_c := int32(sum ^ y)
		t_d := (c.k[int32(p & 3) ^ e]) ^ int32(z)

		v[n-1] += (t_a + t_b) ^ (t_c + t_d)
		z = v[n-1]
	}
}

// blockDecrypt decrypts the []int32 represtentation of a block,
// in-place.
func (c *xxteaCipher) blockDecrypt(v []int32) {
	n := len(v)
	y := v[0]
	z := v[n-1]
	q := 6 + 52/n

	sum := int32(q * delta)
	for sum != 0 {
		e := (sum >> 2) & 3
		var p int
		for p = n - 1; p > 0; p-- {
			z = v[p-1]

			t_a := int32(uint32(z) >> 5) ^ int32(uint32(y) << 2)
			t_b := int32(uint32(y) >> 3) ^ int32(uint32(z) << 4)
			t_c := int32(sum ^ y)
			t_d := (c.k[int32(p & 3) ^ e]) ^ int32(z)

			v[p] -= (t_a + t_b) ^ (t_c + t_d)
			y = v[p]
		}
		z = v[n-1]

		t_a := int32(uint32(z) >> 5) ^ int32(uint32(y) << 2)
		t_b := int32(uint32(y) >> 3) ^ int32(uint32(z) << 4)
		t_c := int32(sum ^ y)
		t_d := (c.k[int32(p & 3) ^ e]) ^ int32(z)
		v[0] -= (t_a + t_b) ^ (t_c + t_d)
		y = v[0]
		sum -= delta
	}
}
