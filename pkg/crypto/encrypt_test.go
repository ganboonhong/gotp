package crypto

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	want := "datatoencrypt123"
	key := "secretkey"
	encrypted := Encrypt(want, key)
	got := Decrypt(encrypted, key)

	if got != want {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}
