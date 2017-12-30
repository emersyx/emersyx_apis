package emircapi

import(
    "errors"
    "plugin"
)

// This function calls the NewIRCOptions function exported by the specified plugin and returns the same value returned
// by the exported function.
func NewIRCOptions(plug *plugin.Plugin) (IRCOptions, error) {
    if plug == nil {
        return nil, errors.New("Invalid plugin handle.")
    }

    f, err := plug.Lookup("NewIRCOptions")
    if err != nil {
        return nil, errors.New("IRC plugin does not have the NewIRCOptions symbol.")
    }

    fc, ok := f.( func() IRCOptions )
    if ok == false {
        return nil, errors.New("The NewIRCOptions function does not have the correct signature.")
    }

    return fc(), nil
}

// This function calls the NewIRCBot function exported by the specified plugin and returns the same value returned
// by the exported function.
func NewIRCBot(plug *plugin.Plugin, options ...func (IRCBot) error) (IRCBot, error) {
    if plug == nil {
        return nil, errors.New("Invalid plugin handle.")
    }

    f, err := plug.Lookup("NewIRCBot")
    if err != nil {
        return nil, errors.New("IRC plugin does not have the NewIRCBot symbol.")
    }

    fc, ok := f.( func(options ...func (IRCBot) error) (IRCBot, error) )
    if ok == false {
        return nil, errors.New("The NewIRCBot function does not have the correct signature.")
    }

    return fc(options...)
}
