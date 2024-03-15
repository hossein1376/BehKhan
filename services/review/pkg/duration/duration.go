package duration

import (
	"encoding/json"
	"time"
)

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(input []byte) error {
	var data string
	if err := json.Unmarshal(input, &data); err != nil {
		return err
	}

	duration, err := time.ParseDuration(data)
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}
