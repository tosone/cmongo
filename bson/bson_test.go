package bson

import "testing"

func TestBsonNew(t *testing.T) {
	var b = New()
	b.AppendInt32("key", 123)
	b.AppendInt64("key", 123)
}
