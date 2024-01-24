package encrypt

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestEncryptPwd(t *testing.T) {
	hash1, err := EncryptPwd("123")
	if err != nil {
		panic(err)
	}
	hash2, err := EncryptPwd("123")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, hash1, hash2)
}
