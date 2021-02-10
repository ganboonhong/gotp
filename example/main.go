package main

import (
	"fmt"

	"github.com/ganboonhong/gotp/pkg/otp"
)

func main() {
	fmt.Println("Random secret:", otp.RandomSecret(16))
	defaultTOTPUsage()
	// defaultHOTPUsage()
}

func defaultTOTPUsage() {
	// otp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	otp := otp.NewDefaultTOTP("MCWFKC6VWWVIDGYC4ZULRKSLQWC7GROF")

	fmt.Println("current one-time password is:", otp.Now())
	// fmt.Println("one-time password of timestamp 0 is:", otp.At(0))
	// fmt.Println(otp.ProvisioningURI("demoAccountName", "issuerName"))

	// fmt.Println(otp.Verify("179394", 1524485781))
}

func defaultHOTPUsage() {
	otp := otp.NewDefaultHOTP("4S62BZNFXXSZLCRO")

	fmt.Println("one-time password of counter 0 is:", otp.At(0))
	fmt.Println(otp.ProvisioningURI("demoAccountName", "issuerName", 1))

	fmt.Println(otp.Verify("944181", 0))
}
