package emtgapi

import (
	"errors"
	"plugin"
)

// NewTelegramOptions calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewTelegramOptions(plug *plugin.Plugin) (TelegramOptions, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewTelegramOptions")
	if err != nil {
		return nil, errors.New("the Telegram plugin does not have the NewTelegramOptions symbol")
	}

	fc, ok := f.(func() TelegramOptions)
	if ok == false {
		return nil, errors.New("the TelegramOptions function does not have the correct signature")
	}

	return fc(), nil
}

// NewTelegramGateway calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewTelegramGateway(plug *plugin.Plugin, options ...func(TelegramGateway) error) (TelegramGateway, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewTelegramGateway")
	if err != nil {
		return nil, errors.New("the Telegram plugin does not have the NewTelegramGateway symbol")
	}

	fc, ok := f.(func(options ...func(TelegramGateway) error) (TelegramGateway, error))
	if ok == false {
		return nil, errors.New("the NewTelegramGateway function does not have the correct signature")
	}

	return fc(options...)
}
