package errs

import "errors"

var (
	ErrorNotCached       = errors.New("not cached")
	ErrorCantCacheMemory = errors.New("cant cache in memory")

	ErrorSteamNotInitialized  = errors.New("steam not initialized")
	ErrorSteamRestartRequired = errors.New("restart required")

	ErrorNotHandled = errors.New("not handled")
)
