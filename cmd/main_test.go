package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	got := Hello()
	assert.Equal(t, want, got, fmt.Sprintf("Hello() = %q, want %q", got, want))
}
