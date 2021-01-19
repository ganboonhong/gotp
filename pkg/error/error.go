package error

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/charmbracelet/glamour"
)

// NoAccount returns a message that indicates database account is empty
func NoAccount() string {

	msg := heredoc.Doc(`
		Please set your database __username__ and __password__ with:
		~~~js
		$ gotp user create
		~~~
	`)
	msg, _ = glamour.Render(msg, "dark")
	return msg
}

// NoAccount returns a message that indicates user has no parameter to generate OTP
func NoParameter() string {

	msg := heredoc.Doc(`
		Please setup your OTP parameter with:
		~~~js
		$ gotp otp create
		~~~
	`)
	msg, _ = glamour.Render(msg, "dark")
	return msg
}
