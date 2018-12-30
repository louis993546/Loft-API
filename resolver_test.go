package loft_test

import (
	"context"
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
		t.Fatalf("Echo time format must be %s, but got %s instead", expectedFormat, echo.Format)
	}

	// TODO: regex for time to make sure it is UTC
}
