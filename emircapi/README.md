# emersyx IRC APIs

## Go plugin API

The interfaces, structs and functions defined in these sources must be followed by IRC resource implementations. A
complete go plugin needs to follow the rules below:

* implement the `IRCBot` interface (defined in file `ircbot.go`)
* implement the `IRCOptions` interface (defined in file `ircopt.go`)
* export a function called `NewIRCBot` with the following signature
```
func NewIRCBot(options ...func (emircapi.IRCBot) error) (emircapi.IRCBot, error)
```
* export a function called `NewIRCOptions` with the following signature
```
func NewIRCOptions() emircapi.IRCOptions
```
* the `IRCBot` implementation needs to make use of the `Message` class (defined in file `message.go`) as the default
  emersyx event type

The two functions `NewIRCBot` and `NewIRCOptions` are required by the wrapper functions with the same names (defined in
file `plugin.go`).

## Using implementations

Implementations of this API are to be distributed as either one go plugin file (e.g. one `.so` file for linux
platforms) or as source code which can be built into one go plugin.

The wrapper functions `NewIRCBot` and `NewIRCOptions` (defined in `plugin.go`) can be used to load the plugin files and
call the exported functions.

## Example implementation

An example implementation of an IRC resource for emersyx can be found in the [emersyx_irc][1] repository.

[1]: https://github.com/emersyx/emersyx_irc
