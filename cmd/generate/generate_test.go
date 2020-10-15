package generate

import (
	"strings"
	"testing"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
)

func TestGenerateTOTP(t *testing.T) {
	f := &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
	}
	chooseType := false

	msg, err := generate(f, chooseType)
	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(msg, "Your OTP: ") {
		t.Error("Failed to generate OTP")
	}
}
