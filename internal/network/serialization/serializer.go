package serialization

import "encoding/json"

func Encode(
	value interface{},
) ([]byte, error) {

	return json.Marshal(value)
}
