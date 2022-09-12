package pkg

import "fmt"

type Env string

const (
	Local   Env = "LOCAL"
	Staging Env = "STAGING"
	Live    Env = "LIVE"
)

var allowedEnv = map[string]Env{
	Local.String():   Local,
	Staging.String(): Staging,
	Live.String():    Live,
}

func NewEnv(str string) (Env, error) {
	if env, ok := allowedEnv[str]; ok {
		return env, nil
	} else {
		return "", fmt.Errorf(
			"%s is not a valid environment value",
			str,
		)
	}
}

func (e Env) String() string {
	return string(e)
}
