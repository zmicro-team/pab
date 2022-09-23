package filecipher

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Encrypt_Decrypt_CBC_TripleDES(t *testing.T) {
	org := []byte("hello world")
	key := make([]byte, 24)
	rand.Read(key)

	cipherText, err := EncryptCBCTripleDES(org, key)
	require.NoError(t, err)

	got, err := DecryptCBCTripleDES(cipherText, key)
	require.NoError(t, err)

	require.Equal(t, org, got)
}
