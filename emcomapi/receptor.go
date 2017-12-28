package emcomapi

// This is the interface for all event receptors part of the emersyx platform. The emersyx core makes use of this type
// when loading all receptors. The event routers make use of this type when receiving events from receptors.
type Receptor interface {
    GetEventsChannel() chan Event
}
