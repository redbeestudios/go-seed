package pkg

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStringFromPathOrDefault(
	key string,
	r *http.Request,
	defaultValue string,
) string {
	if str, err := GetStringFromPath(key, r); err != nil {
		return defaultValue
	} else {
		return str
	}
}

func GetStringFromPath(
	key string,
	r *http.Request,
) (string, error) {
	str := mux.Vars(r)[key]

	if len(str) < 1 {
		return "", fmt.Errorf("No key %s found in path", key)
	}

	return str, nil
}
