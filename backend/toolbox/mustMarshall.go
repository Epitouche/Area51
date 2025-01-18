package toolbox

import "encoding/json"

type Marshal interface {
	MustMarshal(v interface{}) string
	RealObject(v interface{}) []byte
}

func MustMarshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func RealObject(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
