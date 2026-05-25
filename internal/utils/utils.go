package utils

import "encoding/json"

func ToJSONString(model any) (string, error) {
	bytes, err := json.Marshal(model)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
