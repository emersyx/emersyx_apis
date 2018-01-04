package emtgapi

import (
	"emersyx.net/emersyx_apis/emcomapi"
)

// The TelegramResource interface specifies the methods that must be implemented by any supported Telegram resource.
type TelegramResource interface {
	// This method must call the getMe method from the Telegram Bot API. It must return the response formatted into a
	// User instance.
	GetMe() (User, error)
	// This method must call the sendMessage method from the Telegram Bot API. The method must return the response
	// formatted into a Message object.
	SendMessage(params TelegramParameters) (Message, error)
}

// The TelegramBot interface specifies the methods that must be implemented by a complete Telegram receptor and
// resource. The reference implementation at https://github.com/emersyx/emersyx_telegram follows this interface.
type TelegramBot interface {
	emcomapi.Receptor
	TelegramResource
}
