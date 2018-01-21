package emcomapi

import (
	"errors"
	"plugin"
)

// Processor is the interface for all event processors part of the emersyx platform. Each processor component must
// implement the Identifiable interface and implement a method via which new Event objects are received for processing.
type Processor interface {
	Identifiable
	// ProcessEvent must process the Event object it receives. The function can be blocking. The emersyx event router is
	// supposed to call this method withing a goroutine. If the event cannot be processed succesfully, an error must be
	// returned.
	ProcessEvent(ev Event) error
	// LoadConfig must load a configuration file for the specific Processor implementation. If the configuration file
	// cannot be loaded, an error must be returned.
	LoadConfig(path string) error
}

// NewProcessor calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewProcessor(plug *plugin.Plugin, id string, cfg string) (Processor, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewProcessor")
	if err != nil {
		return nil, errors.New("the processor plugin does not have the NewProcessor symbol")
	}

	fc, ok := f.(func(id string, cfg string) (Processor, error))
	if ok == false {
		return nil, errors.New("the NewProcessor function does not have the correct signature")
	}

	return fc(id, cfg)
}
