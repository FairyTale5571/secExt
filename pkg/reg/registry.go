package reg

import (
	"strings"

	"github.com/fairytale5571/secExt/pkg/logger"
	"golang.org/x/sys/windows/registry"
)

type Reg struct {
	logger *logger.Wrapper
}

func New() *Reg {
	return &Reg{
		logger: logger.New("registry"),
	}
}

func (r *Reg) getGoCategory(category string) registry.Key {
	var goCategory registry.Key
	switch strings.ToLower(category) {
	case "classes_root":
		goCategory = registry.CLASSES_ROOT
	case "current_user":
		goCategory = registry.CURRENT_USER
	case "local_machine":
		goCategory = registry.LOCAL_MACHINE
	case "users":
		goCategory = registry.USERS
	case "current_config":
		goCategory = registry.CURRENT_CONFIG
	default:
		// TODO: Logger
	}
	return goCategory
}

func (r *Reg) WriteReg(category, path, key, value string) string {
	goCategory := r.getGoCategory(category)

	k, err := registry.OpenKey(goCategory, path, registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS)
	defer k.Close()

	if err != nil {
		k, _, err = registry.CreateKey(goCategory, path, registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS)
		if err != nil {
			return err.Error()
		}
	}

	err = k.SetStringValue(key, value)
	if err != nil {
		return err.Error()
	}
	return "Written"
}

func (r *Reg) ReadReg(category, path, value string) (string, error) {
	goCategory := r.getGoCategory(category)
	k, err := registry.OpenKey(goCategory, path, registry.QUERY_VALUE|registry.WOW64_64KEY)
	defer k.Close()
	if err != nil {
		return err.Error(), err
	}

	s, _, err := k.GetStringValue(value)
	if err != nil {
		return err.Error(), err
	}
	return s, nil
}

func (r *Reg) DelReg(category, path, value string) string {
	goCategory := r.getGoCategory(category)
	if err := registry.DeleteKey(goCategory, path); err != nil {
		return err.Error()
	}
	return "Deleted"
}
