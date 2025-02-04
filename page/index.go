package page

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type IndexContent struct {
	Text string
}

func Index(c IndexContent) Node {
	return Div(
		Text(c.Text),
	)
}
