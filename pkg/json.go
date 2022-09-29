package pkg

import "encoding/json"

func MustMarshal(elem interface{}) string {
	if res, err := json.Marshal(elem); err != nil {
		panic(err)
	} else {
		return string(res)
	}
}
