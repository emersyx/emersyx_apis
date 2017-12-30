package emtgapi

// The TelegramBot interface specifies the methods that must be implemented by any supported IRC resource. The reference
// implementation of the IRC resource at https://github.com/emersyx/emersyx_telegram follows this interface.
type TelegramBot interface {
    SendPrivateMessage(userID int64, message string) error
    SendChannelMessage(channelName string, message string) error
}
