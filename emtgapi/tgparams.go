package emtgapi

type TgParams interface {
    // Set the chat_id parameter. In the official documentation this parameter is of type Integer or String. If you wish
    // to use an Integer as parameter, then simply pass a go string object with the contents of the Integer.
    ChatID(value string) error
    // Set the text parameter.
    Text(value string) error
    // Set the parse_mode parameter.
    ParseMode(value string) error
    // Set the disable_web_page_preview parameter.
    DisableWebPagePreview(value bool) error
    // Set the disable_notification parameter.
    DisableNotification(value bool) error
    // Set the reply_to_message_id parameter.
    ReplyToMessageID(value string) error
    // Set the reply_markup parameter. In the official documentation, this parameter is of type InlineKeyboardMarkup,
    // ReplyKeyboardMarkup, ReplyKeyboardRemove or ForceReply. Based on the documentation, the parameter is expected to
    // be JSON-serialized. This method takes a string as argument which should be a JSON string of the appropriate
    // object.
    ReplyMarkup(value string) error
}
