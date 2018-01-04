package emircapi

import (
	"errors"
	"plugin"
)

// NewIRCOptions calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewIRCOptions(plug *plugin.Plugin) (IRCOptions, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewIRCOptions")
	if err != nil {
		return nil, errors.New("IRC plugin does not have the NewIRCOptions symbol")
	}

	fc, ok := f.(func() IRCOptions)
	if ok == false {
		return nil, errors.New("the NewIRCOptions function does not have the correct signature")
	}

	return fc(), nil
}

// NewIRCBot calls the function with the same name exported by the specified plugin and returns the same value returned
// by the exported function.
func NewIRCBot(plug *plugin.Plugin, options ...func(IRCBot) error) (IRCBot, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewIRCBot")
	if err != nil {
		return nil, errors.New("IRC plugin does not have the NewIRCBot symbol")
	}

	fc, ok := f.(func(options ...func(IRCBot) error) (IRCBot, error))
	if ok == false {
		return nil, errors.New("the NewIRCBot function does not have the correct signature")
	}

	return fc(options...)
}
