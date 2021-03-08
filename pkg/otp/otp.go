package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"hash"
	"math"
	"strings"

	"github.com/ganboonhong/gotp/pkg/parameter"
)

type Hasher struct {
	HashName string
	Digest   func() hash.Hash
}

type OTP struct {
	secret string  // secret in base32 format
	digits int     // number of integers in the OTP. Some apps expect this to be 6 digits, others support more.
	hasher *Hasher // digest function to use in the HMAC (expected to be sha1)
}

func NewOTP(secret string, digits int, hasher *Hasher) OTP {
	if hasher == nil {
		hasher = &Hasher{
			HashName: parameter.DefaultAlgorithm,
			Digest:   sha1.New,
		}
	}
	return OTP{
		secret: secret,
		digits: digits,
		hasher: hasher,
	}
}

/*
params
    input: the HMAC counter value to use as the OTP input. Usually either the counter, or the computed integer based on the Unix timestamp
*/
func (o *OTP) generateOTP(input int) string {
	if input < 0 {
		panic("input must be positive integer")
	}
	hasher := hmac.New(o.hasher.Digest, o.byteSecret())
	hasher.Write(Itob(input))
	hmacHash := hasher.Sum(nil)

	// https://tools.ietf.org/html/rfc4226
	// 5.4.  Example of HOTP Computation for Digit = 6
	offset := int(hmacHash[len(hmacHash)-1] & 0xf)
	code := ((int(hmacHash[offset]) & 0x7f) << 24) |
		((int(hmacHash[offset+1] & 0xff)) << 16) |
		((int(hmacHash[offset+2] & 0xff)) << 8) |
		(int(hmacHash[offset+3]) & 0xff)

	// We then take this number modulo 1,000,000 (10^6)
	// to generate the 6-digit HOTP value 872921 decimal.
	code = code % int(math.Pow10(o.digits))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", o.digits), code)
}

func (o *OTP) byteSecret() []byte {
	secret := strings.ToUpper(o.secret)
	bytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	return bytes
}
