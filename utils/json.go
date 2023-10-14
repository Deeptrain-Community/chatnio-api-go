package utils

import "encoding/json"

func MarshalForm(data interface{}) string {
	res, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(res)
}

func UnmarshalForm[T any](data string) *T {
	var res T
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil
	}
	return &res
}

func UnmarshalByteForm[T any](data []byte) (*T, error) {
	var res T
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
