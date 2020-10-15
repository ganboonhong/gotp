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
		$ gotp db set
		~~~
	`)
	msg, _ = glamour.Render(msg, "dark")
	return msg
}
