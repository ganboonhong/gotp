package otp

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/stretchr/testify/assert"
)

var totp = NewDefaultTOTP("4S62BZNFXXSZLCRO")

func TestTOTP_Now(t *testing.T) {
	// arrange
	expected := totp.At(CurrentTimestamp())
	err := "TOTP generate otp error!"

	// act
	actual := totp.Now()

	// assert
	assert.Equal(t, expected, actual, err)
}

func TestTOTP_NowWithExpiration(t *testing.T) {
	otp, exp := totp.NowWithExpiration()
	cts := CurrentTimestamp()
	if otp != totp.Now() {
		t.Error("TOTP generate otp error!")
	}
	if totp.At(cts+parameter.DefaultInterval) != totp.At(int(exp)) {
		t.Error("TOTP expiration otp error!")
	}
}

func TestTOTP_Verify(t *testing.T) {
	if !totp.Verify("179394", 1524485781) {
		t.Error("verify failed")
	}
}

func TestTOTP_ProvisioningURI(t *testing.T) {
	expect := "otpauth://totp/github:xlzd?secret=4S62BZNFXXSZLCRO&issuer=github"
	uri := totp.ProvisioningURI("xlzd", "github")
	if expect != uri {
		t.Error("ProvisioningURI error")
	}
}
