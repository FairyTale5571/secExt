package steam

import (
	"fmt"
	"unsafe"

	"github.com/fairytale5571/secExt/pkg/cache"
	"github.com/fairytale5571/secExt/pkg/errs"
	"github.com/hajimehoshi/go-steamworks"
)

const (
	steamAppId = 107410
)

type Steam struct {
	cache *cache.Config
}

func New() (*Steam, error) {
	res := &Steam{
		cache: cache.SetupCache(),
	}

	if steamworks.RestartAppIfNecessary(steamAppId) {
		return nil, errs.ErrorSteamRestartRequired
	}
	if !steamworks.Init() {
		return nil, errs.ErrorSteamNotInitialized
	}
	return res, nil
}

func (s *Steam) GetPlayerUid() string {
	if res, err := s.cache.Get("SteamId"); err != nil && res != "" {
		return res
	}

	sid := steamworks.SteamUser().GetSteamID()
	if unsafe.Sizeof(int(0)) == 4 {
		sid += 76561197960265728
	}
	err := s.cache.Set("SteamId", fmt.Sprintf("%d", sid))
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%d", sid)
}
