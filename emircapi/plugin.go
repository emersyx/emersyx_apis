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
		return nil, errors.New("the IRC plugin does not have the NewIRCOptions symbol")
	}

	fc, ok := f.(func() IRCOptions)
	if ok == false {
		return nil, errors.New("the NewIRCOptions function does not have the correct signature")
	}

	return fc(), nil
}

// NewIRCGateway calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewIRCGateway(plug *plugin.Plugin, options ...func(IRCGateway) error) (IRCGateway, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewIRCGateway")
	if err != nil {
		return nil, errors.New("the IRC plugin does not have the NewIRCGateway symbol")
	}

	fc, ok := f.(func(options ...func(IRCGateway) error) (IRCGateway, error))
	if ok == false {
		return nil, errors.New("the NewIRCGateway function does not have the correct signature")
	}

	return fc(options...)
}
