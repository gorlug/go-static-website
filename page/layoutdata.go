package page

type WebsiteLayoutData struct {
	WebsiteMetaData
	Lang         Language
	NavItemNames NavItemNames
}

type WebsiteMetaData struct {
	Title       string
	Description string
	SelectedNav NavItemName
}

type NavItemName string

const (
	NavItemNameStart NavItemName = "Start"
	NavItemNameAbout NavItemName = "About"
	NavItemNameNone  NavItemName = ""
)

type NavItem struct {
	Label string
	Url   string
	Name  NavItemName
}

type NavItemNames struct {
	Start NavItem
	About NavItem
}

type Language string

const (
	De Language = "de"
	En Language = "en"
)
