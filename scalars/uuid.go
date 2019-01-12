package scalars

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

// MarshalUUIDScalar writes uuid as []byte into graphql
func MarshalUUIDScalar(uuid uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write(uuid.Bytes())
	})
}

// UnmarshalUUIDScalar reads data from fields in graphql that's mark as UUID, and converts it to uuid.UUID.
// Currently it just panic when the input is not in the form of string
func UnmarshalUUIDScalar(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.FromString(v)
	case uuid.UUID:
		return v, nil
	default:
		panic("what should I return when it's a type that I cannot possibly parse?")
	}
}
