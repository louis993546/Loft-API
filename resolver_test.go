package loft_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/louistsaitszho/loft"
)

func TestEcho(t *testing.T) {
	var resolver loft.Resolver
	echo, err := resolver.Query().Echo(context.Background())

	if err != nil {
		t.Fatal("Echo returns error when it's not suppose to")
	}

	expectedFormat := "RFC3339"
	if echo.Format != expectedFormat {
		t.Fatalf("Echo time format must be %s, but got %s instead\n", expectedFormat, echo.Format)
	}

	pattern := "^(\\d{4})(-)(\\d{2})(-)(\\d{2})(T)(\\d{2})(:)(\\d{2})(:)(\\d{2})(Z)$"
	r, _ := regexp.Compile(pattern)

	validTimestampPattern := r.MatchString(echo.Time)
	if validTimestampPattern == false {
		t.Fatalf("Looks like the timestamp is not in the right pattern: %s\n", echo.Time)
	}
}
