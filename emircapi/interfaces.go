package emircapi

import(
    "errors"
    "plugin"
)

// The IRCBot interface specifies the methods that must be implemented by any supported IRC resource. The reference
// implementation of the IRC resource at https://github.com/emersyx/emersyx_irc follows this interface.
type IRCBot interface {
    Connect() error
    Quit()
    Join(channel string)
    Privmsg(destination string, message string)
}
