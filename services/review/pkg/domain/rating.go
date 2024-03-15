package domain

import (
	"encoding/json"
	"errors"
)

type Rating uint8

func (r *Rating) UnmarshalJSON(data []byte) error {
	var rating uint8
	if err := json.Unmarshal(data, &rating); err != nil {
		return err
	}

	if rating <= 0 || rating > 10 {
		return errors.New("rating must be a value between 1 and 10")
	}

	*r = Rating(rating)
	return nil
}
