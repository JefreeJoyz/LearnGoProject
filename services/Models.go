package services

/*
{
    "ip": "54.86.50.139"
}
*/

type GetIP struct {
	Ip string
}

/*
{
    "ip": "54.86.50.139",
    "hostname": "ec2-54-86-50-139.compute-1.amazonaws.com",
    "city": "Ashburn",
    "region": "Virginia",
    "country": "US",
    "loc": "39.0437,-77.4875",
    "org": "AS14618 Amazon.com, Inc.",
    "postal": "20147",
    "timezone": "America/New_York",
    "readme": "https://ipinfo.io/missingauth"
}
*/

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
