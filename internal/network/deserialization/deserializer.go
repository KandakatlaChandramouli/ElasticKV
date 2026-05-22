package deserialization

import "encoding/json"

func Decode(
	data []byte,
	value interface{},
) error {

	return json.Unmarshal(
		data,
		value,
	)
}
