package xxtea_test

import (
	"testing"
	"encoding/hex"
	"github.com/syndeck/xxtea"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestLib(t *testing.T) {
	data := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF,
		0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99}
	key := []byte{0xC0, 0xFF, 0xEE, 0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}

	cipher, err := xxtea.NewCipher(key)
	if err != nil {
		log.Fatal(err)
		return
	}

	cleartext_original := data
	var encrypted = make([]byte, len(cleartext_original))

	cipher.Encrypt(encrypted, cleartext_original)

	var cleartext_result = make([]byte, len(cleartext_original))
	cipher.Decrypt(cleartext_result, encrypted)

	assert.Equal(t, hex.EncodeToString(cleartext_original), hex.EncodeToString(cleartext_result))
}
