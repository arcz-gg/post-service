package utils

import "encoding/json"

func MarshalJson[K comparable](data K) ([]byte, error) {

	return json.Marshal(data);
}

func UnMarshalJson[K comparable](data []byte) (*K, error) {

	var target *K;
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err;
	}

	return target, nil;
}
