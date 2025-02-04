package page

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

const DeIndexUrl = "/de/index"
const EnIndexUrl = "/en/index"
const DefaultLanguage = En

func Create(layoutData WebsiteLayoutData, content Node) Node {
	return HTML5(HTML5Props{
		Title:       layoutData.Title,
		Language:    string(layoutData.Lang),
		Description: layoutData.Description,
		Head: []Node{
			Link(Rel("stylesheet"), Href("/static/styles.css")),
			Script(Type("text/javascript"), Src("/static/scripts/language-select.js")),
			addLanguageSelectScript(layoutData),
		},
		Body: []Node{
			Div(
				ID("center"),
				getNav(layoutData),
				Div(
					ID("content"),
					content,
				),
			),
		},
	})
}

func addLanguageSelectScript(metadata WebsiteLayoutData) Node {
	script := InlineTemplate(`
(function() {
	const langs = {
		de: "$.DeUrl$",
		en: "$.EnUrl$",
	};	
	const currentLang = "$.CurrentLang$";
	const defaultLang = "$.DefaultLang$";
	selectWebsiteLanguage(currentLang, langs, defaultLang);
})();`, struct {
		CurrentLang string
		DeUrl       string
		EnUrl       string
		DefaultLang string
	}{
		CurrentLang: string(metadata.Lang),
		DeUrl:       DeIndexUrl,
		EnUrl:       EnIndexUrl,
		DefaultLang: string(DefaultLanguage),
	})
	return Script(Type("text/javascript"), Raw(script))
}

type Routes map[string]Node

func (r Routes) AddRoutes(newRoutes Routes) Routes {
	for route, content := range newRoutes {
		r[route] = content
	}
	return r
}

func getNav(metadata WebsiteLayoutData) Node {
	return Nav(
		Aria("label", "primary"),
		Ul(
			createNavItem(metadata.NavItemNames.Start, metadata.SelectedNav),
			createNavItem(metadata.NavItemNames.About, metadata.SelectedNav),
			createNavItem(NavItem{
				Label: "EN",
				Url:   EnIndexUrl,
				Name:  NavItemNameNone,
			}, metadata.SelectedNav),
			createNavItem(NavItem{
				Label: "DE",
				Url:   DeIndexUrl,
				Name:  NavItemNameNone,
			}, metadata.SelectedNav),
		),
	)
}

func createNavItem(navItem NavItem, selected NavItemName) Node {
	return Li(
		A(
			Href(navItem.Url+".html"),
			Text(navItem.Label),
			If(navItem.Name == selected, Attr("aria-current", "page")),
			If(navItem.Name == selected, Text(" x")),
		),
	)
}
