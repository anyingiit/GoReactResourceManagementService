package globalVars

import "errors"

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
		return errors.New("global var is already set")
	}
	g.val = val
	g.isSet = true
	return nil
}

func (g *GlobalVars) Get() (interface{}, error) {
	if !g.isSet {
		return nil, errors.New("global var is not set")
	}
	return g.val, nil
}
