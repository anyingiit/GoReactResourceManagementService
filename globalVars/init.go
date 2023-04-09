package globalVars

func InitGlobalVars() {
	ProjectRootPath = newProjectRootPath()
	Db = newDb()
	ProjectConfig = newProjectConfig()
}
