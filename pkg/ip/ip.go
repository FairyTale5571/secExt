package ip

import (
	"fmt"

	goip "github.com/FairyTale5571/go-ip-api"
	"github.com/fairytale5571/secExt/pkg/logger"
	"github.com/rdegges/go-ipify"
)

type IP struct {
	logger *logger.Wrapper
}

func New() *IP {
	return &IP{
		logger: logger.New("ip"),
	}
}

func (i *IP) GetGeoIp() string {
	client := goip.NewClient()
	res, err := client.GetLocationForIp(i.GetIp())
	defer res.Close()
	if err != nil {
		i.logger.Errorf("Cant get geo ip: %s", err)
	}

	return fmt.Sprintf(`["%s","%s","%s","%s","%s","%s"]`,
		res.City,
		res.Country,
		res.CountryCode,
		res.Region,
		res.RegionName,
		res.Zip)
}

func (i *IP) GetIp() string {
	ip, err := ipify.GetIp()
	if err != nil {
		i.logger.Errorf("Cant get ip: %s", err)
		return "Cant get ip"
	}
	return ip
}
