package emcomapi

// This is the interface for all events supported by the various emersyx components. Event routers should use this type
// if they are to support two or more event types. If event routers only support one event type (e.g. emircapi.Message
// or emtgapi.Update), then they can use that particular type directly.
type Event interface {
    GetSourceIdentifier() string
}
