package flexmessage

import "encoding/json"

type Component map[string]any

func ToJSON(c Component) (string, error) {
	data, err := json.MarshalIndent(c, "", "  ")
	return string(data), err
}
