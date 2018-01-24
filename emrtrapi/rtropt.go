package emrtrapi

import (
	"emersyx.net/emersyx_apis/emcomapi"
)

// The RouterOptions interface specifies the options which can be set on a Router implementation. Each method returns a
// function, which can apply the appropriate configuration to the Router implementation. Each Router implementation
// needs to also provide a related RouterOptions implementation. The return values of each method of the RouterOptions
// implementation must be directly usable as arguments to the NewRouter implementation. Different RouterOptions
// implementations may not be compatible with the same Router implementation.
type RouterOptions interface {
	Gateways(gws ...emcomapi.Identifiable) func(Router) error
	Processors(procs ...emcomapi.Processor) func(Router) error
	Routes(routes map[string][]string) func(Router) error
}
