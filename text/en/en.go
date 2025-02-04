package en

import (
	"go-static-website/page"
	. "maragu.dev/gomponents"
)

const urlPrefix = "/en"

func Page(metaData page.WebsiteMetaData, content Node) Node {
	return page.Create(
		page.WebsiteLayoutData{
			WebsiteMetaData: metaData,
			Lang:            page.En,
			NavItemNames:    getNavItemNames(),
		},
		content,
	)
}

func IndexPage() Node {
	return Page(
		page.WebsiteMetaData{
			Title:       "Index",
			Description: "Index page",
			SelectedNav: page.NavItemNameStart,
		},
		page.Index(
			page.IndexContent{
				Text: "Hello World, this is Index",
			},
		),
	)
}

const AboutUrl = urlPrefix + "/about"

func AboutPage() Node {
	return Page(
		page.WebsiteMetaData{
			Title:       "About",
			Description: "About page",
			SelectedNav: page.NavItemNameAbout,
		},
		page.About(page.AboutContent{
			Title: "About",
			Text:  "This is the about page",
			Technologies: page.AboutTechnologies{
				Go:         "Go is the programming language",
				Gomponents: "Gomponents is used to create HTML elements as Go code",
			},
		}),
	)
}

func Routes() page.Routes {
	return page.Routes{
		page.EnIndexUrl: IndexPage(),
		AboutUrl:        AboutPage(),
	}
}

func getNavItemNames() page.NavItemNames {
	return page.NavItemNames{
		Start: page.NavItem{
			Label: "Start",
			Url:   page.EnIndexUrl,
			Name:  page.NavItemNameStart,
		},
		About: page.NavItem{
			Label: "About",
			Url:   AboutUrl,
			Name:  page.NavItemNameAbout,
		},
	}
}
