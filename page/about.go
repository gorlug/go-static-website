package page

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func About(c AboutContent) Node {
	return Div(
		H1(Text(c.Title)),
		P(Text(c.Text)),
		Ul(
			Li(Text(c.Technologies.Go)),
			Li(Text(c.Technologies.Gomponents)),
		),
	)
}

type AboutContent struct {
	Title        string
	Text         string
	Technologies AboutTechnologies
}

type AboutTechnologies struct {
	Go         string
	Gomponents string
}
