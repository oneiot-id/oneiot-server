package helper

import "encoding/json"

func MarshalThis(any interface{}) string {
	marshalled, err := json.Marshal(any)

	if err != nil {
		return ""
	}

	return string(marshalled)
}
