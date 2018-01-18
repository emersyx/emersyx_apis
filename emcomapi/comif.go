package emcomapi

// Event is the interface for all events supported by the various emersyx components. Event routers should use this type
// if they are to support two or more event types. If event routers only support one event type (e.g. emircapi.Message
// or emtgapi.Update), then they can use that particular type directly.
type Event interface {
	// GetSourceIdentifier must return the identifier of the emersyx receptor which generated the event.
	GetSourceIdentifier() string
}

// Identifiable is a low-level interface (w.r.t. hierarchy of types in the emersyx framework) for all components which
// have to be uniquely identifiable. This includes resources, receptors, gateways (i.e. components which are
// simultaneously resources and receptors) and processors.
type Identifiable interface {
	// GetIdentifier must return the identifier of the component.
	GetIdentifier() string
}

// Receptor is the interface for all event receptors part of the emersyx platform. Each receptor component must expose
// a channel via which events are pushed.
type Receptor interface {
	// GetEventsChannel must return the channel via which the Receptor implementation pushes Event objects.
	GetEventsChannel() chan Event
}
