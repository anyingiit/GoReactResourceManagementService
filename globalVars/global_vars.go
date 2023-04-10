package globalVars

import "errors"

// ErrNotSet is returned when the global variable is not set
var ErrNotSet = errors.New("global var is not set")

// var alraedy set error
var ErrAlreadySet = errors.New("global var is already set")

type GlobalVars struct {
	val     interface{}
	isSet   bool
	varName string
}

// newGlobalVars returns a new GlobalVars struct
func newGlobalVars(varName string) *GlobalVars {
	return &GlobalVars{
		varName: varName,
	}
}

func (g *GlobalVars) Set(val interface{}) error {
	if g.isSet {
		return ErrAlreadySet
	}
	g.val = val
	g.isSet = true
	return nil
}

func (g *GlobalVars) Get() (interface{}, error) {
	if !g.isSet {
		return nil, ErrNotSet
	}
	return g.val, nil
}
