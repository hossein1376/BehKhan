package reqID

import (
	"context"
	"crypto/rand"

	"github.com/oklog/ulid/v2"
)

const RequestIDKey = "request_id"

func NewRequestID() (string, error) {
	id, err := ulid.New(ulid.Now(), rand.Reader)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func GetRequestID(c context.Context) (string, bool) {
	id, ok := c.Value(RequestIDKey).(string)
	return id, ok
}
