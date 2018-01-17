# emersyx Telegram APIs

## Go plugin API

The interfaces, structs and functions defined in these sources must be followed by IRC resource implementations. A
complete go plugin needs to follow the rules below:

* implement the `TelegramGateway` interface (defined in file `tggw.go`)
* implement the `TelegramOptions` interface (defined in file `tgopt.go`)
* export a function called `NewTelegramGateway` with the following signature
```
func NewTelegramGateway(options ...func (emtgapi.TelegramGateway) error) (emtgapi.TelegramGateway, error)
```
* export a function called `NewTelegramOptions` with the following signature
```
func NewTelegramOptions() emtgapi.TelegramOptions
```
* the `TelegramGateway` implementation needs to make use of the `Update` class (defined in file `types.go`) as the
  default emersyx event type

The two functions `NewTelegramGateway` and `NewTelegramOptions` are required by the wrapper functions with the same
names (defined in file `plugin.go`).

## Telegram Bot API methods

This API tries to follow the names from the official documentation as much as possible. This makes it easy to explore
the Bot API and this API together.

Setting parameters for the Bot API methods should be done using implementations of the `TelegramParameters` interface.
Implementations of this interface should ensure that each method sets the appropriate parameter. Methods of the
`TelegramGateway` implementation (for Bot API methods that accept parameters) should accept one `TelegramParameters`
object and use the appropriate parameter values.

## Using implementations

Implementations of this API are to be distributed as either one go plugin file (e.g. one `.so` file for linux platforms)
or as source code which can be built into one go plugin.

The wrapper functions `NewTelegramGateway` and `NewTelegramOptions` (defined in `plugin.go`) can be used to load the
plugin files and call the exported functions.

## Example implementation

An example implementation of a Telegram gateway for emersyx can be found in the [emersyx_telegram][1] repository.

[1]: https://github.com/emersyx/emersyx_telegram
