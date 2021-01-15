package cmdutil

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ganboonhong/gotp/pkg/database"
	"github.com/joho/godotenv"
)

type Config struct {
	DbUser     string
	DbPassword string
}

const (
	OtpTypeTotp = "totp"
	OtpTypeHotp = "hotp"
)

/*
Returns the provisioning URI for the OTP; works for either TOTP or HOTP.
This can then be encoded in a QR Code and used to provision the Google Authenticator app.
For module-internal use.
See also:
    https://github.com/google/google-authenticator/wiki/Key-Uri-Format

params:
    otpType：     otp type, must in totp/hotp
    secret:       the hotp/totp secret used to generate the URI
    accountName:  name of the account
    issuerName:   the name of the OTP issuer; this will be the organization title of the OTP entry in Authenticator
    algorithm:    the algorithm used in the OTP generation
    initialCount: starting counter value. Only works for hotp
    digits:       the length of the OTP generated code.
    period:       the number of seconds the OTP generator is set to expire every code.

returns: provisioning uri
*/
func BuildUri(otpType, secret, accountName, issuerName, algorithm string, initialCount, digits, period int) string {
	if otpType != OtpTypeHotp && otpType != OtpTypeTotp {
		panic("otp type error, got " + otpType)
	}

	urlParams := make([]string, 0)
	urlParams = append(urlParams, "secret="+secret)
	if otpType == OtpTypeHotp {
		urlParams = append(urlParams, fmt.Sprintf("counter=%d", initialCount))
	}
	label := url.QueryEscape(accountName)
	if issuerName != "" {
		issuerNameEscape := url.QueryEscape(issuerName)
		label = issuerNameEscape + ":" + label
		urlParams = append(urlParams, "issuer="+issuerNameEscape)
	}
	if algorithm != "" && algorithm != "sha1" {
		urlParams = append(urlParams, "algorithm="+strings.ToUpper(algorithm))
	}
	if digits != 0 && digits != 6 {
		urlParams = append(urlParams, fmt.Sprintf("digits=%d", digits))
	}
	if period != 0 && period != 30 {
		urlParams = append(urlParams, fmt.Sprintf("period=%d", period))
	}
	return fmt.Sprintf("otpauth://%s/%s?%s", otpType, label, strings.Join(urlParams, "&"))
}

// get current timestamp
func CurrentTimestamp() int {
	return int(time.Now().Unix())
}

// integer to byte array
func Itob(integer int) []byte {
	byteArr := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		byteArr[i] = byte(integer & 0xff)
		integer = integer >> 8
	}
	return byteArr
}

// generate a random secret of given length
func RandomSecret(length int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")

	bytes := make([]rune, length)

	for i := range bytes {
		bytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(bytes)
}

func GetConfig() Config {
	godotenv.Load()

	return Config{
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	}
}

func GetConfigTest() Config {
	return Config{
		DbUser:     "johndoe",
		DbPassword: "Pwd1234!!",
	}
}

type Factory struct {
	GetConfig func() Config
	DB        *database.DB
}

func NewFactory() *Factory {
	return &Factory{
		GetConfig,
		database.NewDB(nil),
	}
}
