package emtgapi

import(
    "emersyx.net/emersyx_apis/emcomapi"
)

// The TelegramResource interface specifies the methods that must be implemented by any supported Telegram resource.
type TelegramResource interface {
    // This method creates a new object that implements the TgParams interface. With this object, parameters for methods
    // can be generated.
    NewTgParams() TgParams
    // This method must call the getMe method from the Telegram Bot API. It must return the response formatted into a
    // User instance.
    GetMe() User
    // This method must call the sendMessage method from the Telegram Bot API. Based on the documentation, the
    // destination field is either an integer or a string, but the chat_id argument to this method is of type string.
    // For integer values, just use a string with the contents of the integer. The method must return the response
    // formatted into a
    SendMessage(params TgParams) Message
}

// The TelegramBot interface specifies the methods that must be implemented by a complete Telegram receptor and
// resource. The reference implementation at https://github.com/emersyx/emersyx_telegram follows this interface.
type TelegramBot interface {
    emcomapi.Receptor
    TelegramResource
}
