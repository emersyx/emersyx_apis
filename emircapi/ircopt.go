package emircapi

import (
	"io"
)

// The IRCOptions interface specifies the options which can be set on an IRCGateway implementation. Each method returns
// a function, which can apply the appropriate configuration to the IRCGateway implementation. Each IRCGateway
// implementation needs to also provide a related IRCOptions implementation. The return values of each method of the
// IRCOptions implementation must be directly usable as arguments to the NewIRCGateway implementation. Different
// IRCOptions implementations may not be compatible with the same IRCGateway implementation.
type IRCOptions interface {
	Logging(writer io.Writer, level uint) func(IRCGateway) error
	Identifier(id string) func(IRCGateway) error
	Nick(nick string) func(IRCGateway) error
	Ident(ident string) func(IRCGateway) error
	Name(name string) func(IRCGateway) error
	Version(version string) func(IRCGateway) error
	Server(address string, port uint, useSSL bool) func(IRCGateway) error
	QuitMessage(message string) func(IRCGateway) error
}
