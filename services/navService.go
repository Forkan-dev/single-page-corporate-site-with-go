package services

type DomainPart struct {
	PartOne string
	PartTwo string
}
type Nav struct {
	Name string
	Url  string
}
type NavData struct {
	SiteName DomainPart
	Navs     []Nav
}

func GetNavs() NavData {
	return NavData{
		SiteName: DomainPart{
			PartOne: "list",
			PartTwo: "race",
		},
		Navs: []Nav{
			{Name: "home", Url: "/"},
			{Name: "about", Url: "/about"},
			{Name: "portfolio", Url: "/portfolio"},
			{Name: "contact", Url: "/contact"},
		},
	}

}
