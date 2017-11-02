package password

import (
	"encoding/json"
)

func UnMarshal(data []byte) Password {
	password := &Password{}

	err := json.Unmarshal(data, password)

	if err != nil {
		panic(err)
	}

	return *password
}
