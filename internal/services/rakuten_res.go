package services

import (
	"encoding/json"
	"fmt"
)

// by JSON-to-GO (https://mholt.github.io/json-to-go/)
type DecResBodyType struct{}

func DecordResponse(resBody []byte) (DecResBodyType, error) {
	decResBody := DecResBodyType{}
	if err := json.Unmarshal(resBody, &decResBody); err != nil {
		return DecResBodyType{}, fmt.Errorf("While unmarshalling the response json: %w", err)
	}
	return decResBody, nil
}
