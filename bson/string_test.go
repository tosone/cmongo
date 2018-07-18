package bson

import (
	"fmt"
	"testing"
)

func TestStringNew(t *testing.T) {
	s := StringNew("hello")
	fmt.Printf("free: %s\n", s.Free(false))
}
