package emircapi

import(
    "errors"
    "plugin"
)

// This function opens the go plugin at the specified path, calls the NewIRCOptions function exported by the plugin and
// returns the same value returned by the exported function.
func NewIRCOptions(path string) IRCOptions {
    plug, err := plugin.Open(path)
    if err != nil {
        return nil, err
    }
    return NewIRCOptions(plug), nil
}

// This function calls the NewIRCOptions function exported by the specified plugin and returns the same value returned
// by the exported function.
func NewIRCOptions(plug *plugin.Plugin) IRCOptions {
    if plug == nil {
        return nil, errors.New("Invalid plugin handle.")
    }

    f, err := irc_plugin.Lookup("NewIRCOptions")
    if err != nil {
        return nil, errors.New("IRC plugin does not have the NewIRCOptions symbol.")
    }

    fc, ok := f.( func() IRCOptions )
    if ok == false {
        return nil, errors.New("The NewIRCOptions function does not have the correct signature.")
    }

    bot := fc(options...)
    return bot, nil
}

// This function opens the go plugin at the specified path, calls the NewIRCBot function exported by the plugin and
// returns the same value returned by the exported function.
func NewIRCBot(path string, options ...func (IRCBot) error) (IRCBot, error) {
    plug, err := plugin.Open(path)
    if err != nil {
        return nil, err
    }
    return NewIRCBot(plug, options...), nil
}

// This function calls the NewIRCBot function exported by the specified plugin and returns the same value returned
// by the exported function.
func NewIRCBot(plug *plugin.Plugin, options ...func (IRCBot) error) (IRCBot, error) {
    if plug == nil {
        return nil, errors.New("Invalid plugin handle.")
    }

    f, err := irc_plugin.Lookup("NewIRCBot")
    if err != nil {
        return nil, errors.New("IRC plugin does not have the NewIRCBot symbol.")
    }

    fc, ok := f.( func(options ...func (IRCBot) error) (IRCBot, error) )
    if ok == false {
        return nil, errors.New("The NewIRCBot function does not have the correct signature.")
    }

    bot := fc(options...)
    return bot, nil
}
