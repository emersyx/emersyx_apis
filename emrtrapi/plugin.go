package emrtrapi

import (
	"errors"
	"plugin"
)

// NewRouterOptions calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewRouterOptions(plug *plugin.Plugin) (RouterOptions, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewRouterOptions")
	if err != nil {
		return nil, errors.New("the router plugin does not have the NewRouterOptions symbol")
	}

	fc, ok := f.(func() (RouterOptions, error))
	if ok == false {
		return nil, errors.New("the NewRouterOptions function does not have the correct signature")
	}

	return fc()
}

// NewRouter calls the function with the same name exported by the specified plugin and returns the same value returned
// by the exported function.
func NewRouter(plug *plugin.Plugin, options ...func(Router) error) (Router, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewRouter")
	if err != nil {
		return nil, errors.New("the router plugin does not have the NewRouter symbol")
	}

	fc, ok := f.(func(optiosn ...func(Router) error) (Router, error))
	if ok == false {
		return nil, errors.New("the NewRouter function does not have the correct signature")
	}

	return fc()
}
