package services

type GetIP struct {
	Ip string
}

type GetDetailsIP struct {
	Ip       string
	Hostname string
	City     string
	Region   string
	Country  string
	Loc      string
	Org      string
	Postal   string
	Timezone string
	Readme   string
}
