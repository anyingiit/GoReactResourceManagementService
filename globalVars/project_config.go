package globalVars

import (
	"github.com/anyingiit/GoReactResourceManagement/structs"
)

var ProjectConfig *projectConfig

type projectConfig struct {
	*GlobalVars
}

// new projectConfig
func newProjectConfig() *projectConfig {
	return &projectConfig{
		GlobalVars: newGlobalVars("ProjectConfig"),
	}
}

func (p *projectConfig) Set(newVal *structs.ProjectConfig) error {
	return p.GlobalVars.Set(newVal)
}

func (p *projectConfig) Get() (*structs.ProjectConfig, error) {
	val, err := p.GlobalVars.Get()
	if err != nil {
		return nil, err
	}
	return val.(*structs.ProjectConfig), nil
}
