package scalars

import (
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalTimestampScalar writes time.Time as string in RFC3339 format to io.Writer
func MarshalTimestampScalar(timestamp time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(timestamp.Format(time.RFC3339)))
	})
}

// UnmarshalTimestampScalar tries to parse string into RFC3339 time.Time
func UnmarshalTimestampScalar(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	default:
		panic("what should I return when it's a type that I cannot possibly parse?")
	}
}
