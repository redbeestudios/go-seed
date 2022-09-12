package pkg

import (
	"fmt"
	"strings"
)

type Env string

const (
	LocalEnvironment   Env = "LOCAL"
	StagingEnvironment Env = "STAGING"
	LiveEnvironment    Env = "LIVE"
)

var allowedEnv = map[string]Env{
	LocalEnvironment.String():   LocalEnvironment,
	StagingEnvironment.String(): StagingEnvironment,
	LiveEnvironment.String():    LiveEnvironment,
}

func NewEnv(str string) (Env, error) {
	if env, ok := allowedEnv[strings.ToUpper(str)]; ok {
		return env, nil
	}

	return "", fmt.Errorf(
		"%s is not a valid environment value",
		str,
	)
}

func (e Env) String() string {
	return string(e)
}
