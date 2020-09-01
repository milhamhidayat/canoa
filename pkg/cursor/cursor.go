package cursor

import (
	"encoding/base64"
	"time"
)

// Encode will encode cursor from time
func Encode(t time.Time) string {
	timeString := t.Format(time.RFC3339Nano)
	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

// Decode will decode encoded time cursor
func Decode(encodedTime string) (t time.Time, err error) {
	decoded, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(decoded)
	t, err = time.Parse(time.RFC3339Nano, timeString)
	return
}
