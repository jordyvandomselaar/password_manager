package password

import (
	"encoding/json"
)

func Marshal(data interface{}) []byte {
	marshaledData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return marshaledData
}
