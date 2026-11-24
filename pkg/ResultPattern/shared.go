package ResultPattern

import "uuid"

func newTraceID() string {
	return uuid.NewV4().String()
}