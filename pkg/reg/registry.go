package reg

import (
	"strings"

	"golang.org/x/sys/windows/registry"
)

type Reg struct{}

func New() *Reg {
	return &Reg{}
}

func (r *Reg) getGoCategory(category string) registry.Key {
	switch strings.ToLower(category) {
	case "classes_root":
		return registry.CLASSES_ROOT
	case "current_user":
		return registry.CURRENT_USER
	case "local_machine":
		return registry.LOCAL_MACHINE
	case "users":
		return registry.USERS
	case "current_config":
		return registry.CURRENT_CONFIG
	}
	return registry.CURRENT_CONFIG
}

func (r *Reg) WriteReg(category, path, key, value string) error {
	goCategory := r.getGoCategory(category)

	k, err := registry.OpenKey(goCategory, path, registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS)
	defer k.Close()

	if err != nil {
		k, _, err = registry.CreateKey(goCategory, path, registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS)
		if err != nil {
			return err
		}
	}

	err = k.SetStringValue(key, value)
	if err != nil {
		return err
	}
	return nil
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

func (r *Reg) DeleteReg(category, path, value string) error {
	goCategory := r.getGoCategory(category)
	if err := registry.DeleteKey(goCategory, path); err != nil {
		return err
	}
	return nil
}
