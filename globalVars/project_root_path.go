package globalVars

var ProjectRootPath *projectRootPath

type projectRootPath struct {
	*GlobalVars
}

// new projectRootPath
func newProjectRootPath() *projectRootPath {
	return &projectRootPath{
		GlobalVars: newGlobalVars("ProjectRootPath"),
	}
}

func (p *projectRootPath) Set(newVal string) error {
	return p.GlobalVars.Set(newVal)
}

func (p *projectRootPath) Get() (string, error) {
	val, err := p.GlobalVars.Get()
	if err != nil {
		return "", err
	}
	return val.(string), nil
}

// var projectRootPath = struct {
// 	val   string
// 	isSet bool
// }{
// 	val:   "",
// 	isSet: false,
// }

// func SetProjectRootPath(newVal string) error {
// 	if projectRootPath.isSet {
// 		return errors.New("projectRootPath is already set")
// 	}
// 	projectRootPath.val = newVal
// 	projectRootPath.isSet = true
// 	return nil
// }

// func GetProjectRootPath() (string, error) {
// 	if !projectRootPath.isSet {
// 		return "", errors.New("projectRootPath is not set")
// 	}
// 	return projectRootPath.val, nil
// }
