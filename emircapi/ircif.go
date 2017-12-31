package emircapi

import(
    "emersyx.net/emersyx_apis/emcomapi"
)

// The IRCResource interface specifies the methods that must be implemented by any supported IRC resource.
type IRCResource interface {
    Connect() error
    IsConnected() bool
    Quit() error
    Join(channel string) error
    Privmsg(destination string, message string) error
}

// The IRCBot interface specifies the methods that must be implemented by a complete IRC receptor and resource. The
// reference implementation at https://github.com/emersyx/emersyx_irc follows this interface.
type IRCBot interface {
    emcomapi.Receptor
    IRCResource
}
