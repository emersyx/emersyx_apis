package emircapi

// The IRCBot interface specifies the methods that must be implemented by any supported IRC resource. The reference
// implementation of the IRC resource at https://github.com/emersyx/emersyx_irc follows this interface.
type IRCBot interface {
    Connect() error
    IsConnected() bool
    Quit() error
    Join(channel string) error
    Privmsg(destination string, message string) error
}
