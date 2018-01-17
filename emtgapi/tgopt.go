package emtgapi

// The TelegramOptions interface specifies the options which can be set on a TelegramGateway implementation. Each method
// returns a function, which can apply the appropriate configuration to the TelegramGateway implementation. Each
// TelegramGateway implementation needs to also provide a related TelegramOptions implementation. The return values of
// each method of the TelegramOptions implementation must be directly usable as arguments to the NewTelegramGateway
// implementation.  Different TelegramOptions implementations may not be compatible with the same TelegramGateway
// implementation.
type TelegramOptions interface {
	Identifier(id string) func(TelegramGateway) error
	APIToken(token string) func(TelegramGateway) error
	UpdatesLimit(limit uint) func(TelegramGateway) error
	UpdatesTimeout(seconds uint) func(TelegramGateway) error
	UpdatesAllowed(types ...string) func(TelegramGateway) error
}
