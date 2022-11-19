package ds

import (
	"encoding/json"
	"fmt"

	"github.com/fairytale5571/secExt/pkg/cache"
	"github.com/fairytale5571/secExt/pkg/logger"

	discord "github.com/SilverCory/golang_discord_rpc"
)

const (
	AppID = "591067147900289029"
)

type DS struct {
	rpc    *discord.RPCConnection
	logger *logger.Wrapper
	cache  *cache.Config
}

func New() (*DS, error) {
	log := logger.New("discord_rpc")
	rpc := discord.NewRPCConnection(AppID)
	err := rpc.Open()
	if err != nil {
		log.Infof("Error opening connection: %s", err.Error())
		return nil, err
	}

	return &DS{
		rpc:    rpc,
		logger: log,
		cache:  cache.SetupCache(),
	}, nil
}

func (d *DS) parseData() (string, error) {
	if d.cache.IsExist("DiscordId") {
		return d.cache.Get("DiscordId")
	}

	read, err := d.rpc.Read()
	if err != nil {
		return "unknown", err
	}
	var str string
	for _, ch := range read {
		if ch == 0 {
			continue
		}
		str += string(ch)
	}
	str = fmt.Sprint("\n", str)

	if err := d.cache.Set("DiscordId", str); err != nil {
		d.logger.Infof("Error setting cache: %s", err.Error())
	}

	return str, nil
}

func (d *DS) GetID() (string, error) {
	str, err := d.parseData()
	if err != nil {
		return "unknown", err
	}
	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(str), &resp); err != nil {
		return "unknown", err
	}

	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		return "unknown", nil
	}
	user, ok := data["user"].(map[string]interface{})
	if !ok {
		return "unknown", nil
	}
	return fmt.Sprintf("%s", user["id"]), nil
}

func (d *DS) GetUsername() (string, error) {
	str, err := d.parseData()
	if err != nil {
		return "unknown", err
	}
	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(str), &resp); err != nil {
		return "unknown", err
	}

	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		return "unknown", nil
	}
	user, ok := data["user"].(map[string]interface{})
	if !ok {
		return "unknown", nil
	}
	return fmt.Sprintf("%s#%s", user["username"], user["discriminator"]), nil
}
