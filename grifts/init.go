package grifts

import (
	"github.com/gobuffalo/buffalo"
	"zippyrocket.com/www/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
