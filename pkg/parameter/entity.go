package parameter

const (
	Table            = "parameters"
	DefaultAlgorithm = "sha1"
	DefaultDigits    = 6
	// DefaultInterval is the period parameter defines a period that a TOTP code will be valid for, in seconds.
	// ref: https://github.com/google/google-authenticator/wiki/Key-Uri-Format#period
	DefaultInterval = 30
)

type Parameter struct {
	UserId  uint
	Secret  string
	Issuer  string
	Account string
}
