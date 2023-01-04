package ds

import (
	"encoding/json"
	"fmt"
	"time"

	discord "github.com/SilverCory/golang_discord_rpc"
	"github.com/fairytale5571/secExt/pkg/cache"
)

const (
	AppID = "591067147900289029"
)

var c *cache.Config

func init() {
	c = cache.SetupCache()
}

func GetDsName() string {
	if c.IsExist("discord_info") {
		if res, err := c.Get("discord_info"); err == nil {
			return res
		}
	}

	// объявляем канал с типом string
	resCh := make(chan string)

	win := discord.NewRPCConnection(AppID)
	err := win.Open()
	if err != nil {
		return err.Error()
	}

	// создаем таймер с таймаутом в 5 секунд
	timer := time.NewTimer(5 * time.Second)

	// запускаем горутину, которая будет читать ответ
	go func() {
		_str, err := win.Read()
		if err != nil {
			return
		}
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
			return
		}

		data := resp["data"].(map[string]interface{})
		user := data["user"].(map[string]interface{})

		// отправляем результат в канал
		resCh <- fmt.Sprintf(`["%s#%s","%s"]`, user["username"], user["discriminator"], user["id"])
	}()

	// ждем таймер или ответ от горутины
	select {
	case <-timer.C:
		// таймер отработал, возвращаем ошибку
		return "timeout error"
	case res := <-resCh:
		// получили ответ от горутины, возвращаем его
		c.Set("discord_info", res)
		return res
	}
}
