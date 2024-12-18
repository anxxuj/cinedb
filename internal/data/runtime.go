package data

import (
	"fmt"
	"strconv"
)

// Represents movie runtime
type Runtime int32

// Return the JSON-encoded value for the movie runtime in the format "<runtime> mins"
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
