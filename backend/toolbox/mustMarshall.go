package toolbox

import "encoding/json"

type Marshal interface {
	MustMarshal(v interface{}) string
}

func MustMarshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}