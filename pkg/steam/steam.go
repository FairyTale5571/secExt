package steam

import (
	"fmt"

	"github.com/fairytale5571/secExt/pkg/cache"
	"github.com/fairytale5571/secExt/pkg/errs"
	"github.com/hajimehoshi/go-steamworks"
)

const (
	steamAppId = 107410
)

type Steam struct {
	cache  *cache.Config
	inited bool
}

func New() (*Steam, error) {
	res := &Steam{
		cache: cache.SetupCache(),
	}

	return res, res.Init()
}

func (s *Steam) Init() error {
	if steamworks.RestartAppIfNecessary(steamAppId) {
		fmt.Println("need restart")
		return errs.ErrorSteamRestartRequired
	}
	s.inited = steamworks.Init()
	return nil
}

func (s *Steam) GetPlayerUid() string {
	s.Init()
	sid := steamworks.SteamUser().GetSteamID()
	err := s.cache.Set("SteamId", fmt.Sprintf("%d", sid))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%d", 234341)
}

func (s *Steam) GetInstalledPathGame() string {
	s.Init()
	if res, err := s.cache.Get("SteamName"); err != nil && res != "" {
		return res
	}

	name := steamworks.SteamApps().GetAppInstallDir(steamAppId)
	err := s.cache.Set("SteamName", name)
	if err != nil {
		return err.Error()
	}
	return name
}
