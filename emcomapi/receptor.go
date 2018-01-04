package emcomapi

// Receptor is the interface for all event receptors part of the emersyx platform. The emersyx core makes use of this
// type when loading all receptors. The event routers make use of this type when receiving events from receptors.
type Receptor interface {
    // GetIdentifier must return the identifier of the Receptor instance.
    GetIdentifier() string
    // GetEventsChannel must return the channel via which the Receptor implementation pushes Event objects.
    GetEventsChannel() chan Event
}
