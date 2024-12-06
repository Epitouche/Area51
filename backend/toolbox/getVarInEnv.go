package toolbox

import "os"

type Toolbox interface {
	GetInEnv(varWanted string) (envVar string)
}

func GetInEnv(varWanted string) (envVar string) {
	envVar = os.Getenv(varWanted)
	if envVar == "" {
		panic(varWanted + " is not set")
	}
	return envVar
}
