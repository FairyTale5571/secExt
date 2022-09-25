package ds

import (
	"encoding/json"
	"fmt"

	discord "github.com/SilverCory/golang_discord_rpc"
)

const (
	AppID = "591067147900289029"
)

func GetDiscord() string {
	win := discord.NewRPCConnection("591067147900289029")
	err := win.Open()
	if err != nil {
		return err.Error()
	}

	_str, _ := win.Read()
	str := ""
	for _, ch := range _str {
		if ch == 0 {
			continue
		}
		str += string(ch)
	}
	str = fmt.Sprint("\n", str)

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(str), &resp); err != nil {
		return err.Error()
	}

	data := resp["data"].(map[string]interface{})
	user := data["user"].(map[string]interface{})
	return fmt.Sprintf(`["%s#%s","%s"]`, user["username"], user["discriminator"], user["id"])
}
