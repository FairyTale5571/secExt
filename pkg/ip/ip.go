package ip

import (
	"fmt"

	goip "github.com/FairyTale5571/go-ip-api"
	"github.com/rdegges/go-ipify"
)

func GetGeoIp() string {
	client := goip.NewClient()
	res, err := client.GetLocationForIp(GetIp())
	defer res.Close()
	if err != nil {
		return "[]"
	}

	return fmt.Sprintf(`["%s","%s","%s","%s","%s","%s"]`,
		res.City,
		res.Country,
		res.CountryCode,
		res.Region,
		res.RegionName,
		res.Zip)
}

func GetIp() string {
	ip, err := ipify.GetIp()
	if err != nil {
		return "Cant get ip"
	}
	return ip
}
