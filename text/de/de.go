package de

import (
	"go-static-website/page"
	. "maragu.dev/gomponents"
)

const urlPrefix = "/de"

func Page(metaData page.WebsiteMetaData, content Node) Node {
	return page.Create(
		page.WebsiteLayoutData{
			WebsiteMetaData: metaData,
			Lang:            page.De,
			NavItemNames:    getNavItemNames(),
		},
		content,
	)
}

func IndexPage() Node {
	return Page(
		page.WebsiteMetaData{
			Title:       "Index - DE",
			Description: "Index Seite",
			SelectedNav: page.NavItemNameStart,
		},
		page.Index(
			page.IndexContent{
				Text: "Hallo Welt, hier ist der Index",
			},
		))
}

const AboutUrl = urlPrefix + "/info"

func AboutPage() Node {
	return Page(
		page.WebsiteMetaData{
			Title:       "Info",
			Description: "Info Seite",
			SelectedNav: page.NavItemNameAbout,
		},
		page.About(page.AboutContent{
			Title: "Info",
			Text:  "Dies ist die Info Seite",
			Technologies: page.AboutTechnologies{
				Go:         "Go ist die Programmiersprache",
				Gomponents: "Gomponents wird verwendet um HTML Elemente als Go Code zu erstellen",
			},
		}),
	)
}

func Routes() page.Routes {
	return page.Routes{
		page.DeIndexUrl: IndexPage(),
		AboutUrl:        AboutPage(),
	}
}

func getNavItemNames() page.NavItemNames {
	return page.NavItemNames{
		Start: page.NavItem{
			Label: "Start",
			Url:   page.DeIndexUrl,
			Name:  page.NavItemNameStart,
		},
		About: page.NavItem{
			Label: "Info",
			Url:   AboutUrl,
			Name:  page.NavItemNameAbout,
		},
	}
}
