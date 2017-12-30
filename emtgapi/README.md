# emersyx Telegram APIs

## Go plugin API

The interfaces, structs and functions defined in these sources must be followed by IRC resource implementations. A
complete go plugin needs to follow the rules below:

* implement the `TelegramBot` interface (defined in file `tgbot.go`)
* implement the `TelegramOptions` interface (defined in file `tgopt.go`)
* export a function called `NewTelegramBot` with the following signature
```
func NewTelegramBot(options ...func (emtgapi.TelegramBot) error) (emtgapi.TelegramBot, error)
```
* export a function called `NewTelegramOptions` with the following signature
```
func NewTelegramOptions() emtgapi.TelegramOptions
```
* the `TelegramBot` implementation needs to make use of the `Update` class (defined in file `types.go`) as the default
  emersyx event type

The two functions `NewTelegramBot` and `NewTelegramOptions` are required by the wrapper functions with the same names
(defined in file `plugin.go`).

## Using implementations

Implementations of this API are to be distributed as either one go plugin file (e.g. one `.so` file for linux
platforms) or as source code which can be built into one go plugin.

The wrapper functions `NewTelegramBot` and `NewTelegramOptions` (defined in `plugin.go`) can be used to load the plugin
files and call the exported functions.

## Example implementation

An example implementation of a Telegram resource for emersyx can be found in the [emersyx_telegram][1] repository.

[1]: https://github.com/emersyx/emersyx_telegram
