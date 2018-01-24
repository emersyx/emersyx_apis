package emrtrapi

import (
	"emersyx.net/emersyx_apis/emcomapi"
)

// Router is the interface which must be implemented by all emersyx routers. The standard router implementation follows
// this interface. The emersyx core also expects implementations to follow it.
type Router interface {
	// GetGateway must search for the emcomapi.Identifiable object loaded with the Gateways option and with the
	// specified identifier. If such an object is found, then it must be returned.
	GetGateway(id string) (emcomapi.Identifiable, error)
	// Run must start a loop in which events coming from 1a) the emcomapi.Identifiable objects 2a) loaded via the
	// Gateways option and 3a) which also implement the emcomapi.Receptor interface are routed to 1b) the
	// emcomapi.Processor objects 2b) loaded via the Processors option. The events should follow the routes specified
	// via the Routes option. The mentioned options refer to the ones in the emrtrapi.RouterOptions interface.
	Run() error
}
