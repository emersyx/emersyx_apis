package emircapi

// The IRCOptions interface specifies the options which can be set on an IRCBot implementation. Each method returns a
// function, which can apply the appropriate configuration to the IRCBot implementation. Each IRCBot implementation
// needs to also provide a related IRCOptions implementation. The return values of each method of the IRCOptions
// implementation must be directly usable as arguments to the NewIRCBot implementation.  Different IRCOptions
// implementations may not be compatible with the same IRCBot implementation.
type IRCOptions interface {
    Identifier(id string)                           func(IRCBot) error;
    Nick(nick string)                               func(IRCBot) error;
    Ident(ident string)                             func(IRCBot) error;
    Name(name string)                               func(IRCBot) error;
    Version(version string)                         func(IRCBot) error;
    Server(address string, port uint, useSSL bool)  func(IRCBot) error;
    QuitMessage(message string)                     func(IRCBot) error;
}
