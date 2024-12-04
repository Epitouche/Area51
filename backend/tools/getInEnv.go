package tools

import "os"

type EnvVar interface {
	GetInEnv(varWanted string) (string)
}

func GetInEnv(varWanted string) (envVar string) {
	envVar = os.Getenv(varWanted)
	if envVar == "" {
		panic(varWanted + " is not set")
	}
	return envVar
}
