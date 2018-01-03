package emtgapi

// The TelegramOptions interface specifies the options which can be set on a TelegramBot implementation. Each method
// returns a function, which can apply the appropriate configuration to the TelegramBot implementation. Each TelegramBot
// implementation needs to also provide a related TelegramOptions implementation. The return values of each method of
// the TelegramOptions implementation must be directly usable as arguments to the NewTelegramBot implementation.
// Different TelegramOptions implementations may not be compatible with the same TelegramBot implementation.
type TelegramOptions interface {
    Identifier(id string) func(TelegramBot) error
    APIToken(token string) func(TelegramBot) error
    UpdatesLimit(limit uint) func(TelegramBot) error
    UpdatesTimeout(seconds uint) func(TelegramBot) error
    UpdatesAllowed(types ...string) func(TelegramBot) error
}
