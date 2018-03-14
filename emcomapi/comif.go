package emcomapi

import (
	"io"
)

// Event is the interface for all events supported by the various emersyx components. Event routers should use this type
// if they are to support two or more event types. If event routers only support one event type (e.g. emircapi.Message
// or emtgapi.Update), then they can use that particular type directly.
type Event interface {
	// GetSourceIdentifier must return the identifier of the emersyx receptor which generated the event.
	GetSourceIdentifier() string
}

// Identifiable is a low-level interface (w.r.t. hierarchy of types in the emersyx framework) for all components which
// have to be uniquely identifiable. This includes gateways (components which are either resources, receptors or both
// simultaneously) and processors.
type Identifiable interface {
	// GetIdentifier must return the identifier of the component.
	GetIdentifier() string
}

// Receptor is the interface for all event receptors part of the emersyx platform. Each receptor component must expose
// a channel via which events are pushed.
type Receptor interface {
	// GetEventsChannel must return the channel via which the Receptor implementation pushes Event objects. The channel
	// is read-only and can not be written to.
	GetEventsChannel() <-chan Event
}

// The RouterOptions interface specifies the options which can be set on a Router implementation. Each method returns a
// function, which can apply the appropriate configuration to the Router implementation. Each Router implementation
// needs to also provide a related RouterOptions implementation. The return values of each method of the RouterOptions
// implementation must be directly usable as arguments to the Router.SetOptions implementation. Different RouterOptions
// implementations may not be compatible with the same Router implementation.
type RouterOptions interface {
	Gateways(gws ...Identifiable) func(Router) error
	Processors(procs ...Processor) func(Router) error
	Routes(routes map[string][]string) func(Router) error
}

// Router is the interface which must be implemented by all emersyx routers. The standard router implementation follows
// this interface. The emersyx core also expects implementations to follow it.
type Router interface {
	// SetOptions sets the router options given as arguments. These are options implemented via RouterOptions.
	SetOptions(options ...func(Router) error) error
	// GetGateway must search for the Identifiable object loaded with the Gateways option and with the specified
	// identifier. If such an object is found, then it must be returned.
	GetGateway(id string) (Identifiable, error)
	// Run must start a loop in which events coming from 1a) the Identifiable objects 2a) loaded via the Gateways option
	// and 3a) which also implement the Receptor interface are routed to 1b) the Processor objects 2b) loaded via the
	// Processors option. The events should follow the routes specified via the Routes option. The mentioned options
	// refer to the ones in the emcomapi.RouterOptions interface.
	Run() error
}

// The ProcessorOptions interface specifies the options which can be set on a Processor implementation. Each method
// returns a function, which can apply the appropriate configuration to the Processor implementation. Each Processor
// implementation needs to also provide a related ProcessorOptions implementation. The return values of each method of
// the ProcessorOptions implementation must be directly usable as arguments to the NewProcessor implementation.
// Different ProcessorOptions implementations may not be compatible with the same Processor implementation.
type ProcessorOptions interface {
	Logging(writer io.Writer, level uint) func(Processor) error
	Identifier(id string) func(Processor) error
	Config(cfg string) func(Processor) error
	Router(rtr Router) func(Processor) error
}

// Processor is the interface for all event processors part of the emersyx platform. Each processor component must
// implement the Identifiable interface and implement a method via which new Event objects are received for processing.
type Processor interface {
	Identifiable
	// GetInEventsChannel must return the channel via which the Processor implementation receives Event objects. The
	// channel is write-only and can not be read from.
	GetInEventsChannel() chan<- Event
	// GetOutEventsChannel must return the channel via which the Processor implementation pushes Event objects. The
	// channel is read-only and can not be written to.
	GetOutEventsChannel() <-chan Event
}
