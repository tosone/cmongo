package bson

import "testing"

func TestNewContext(t *testing.T) {
	NewContext(ContextNone)
	NewContext(ContextThreadSafe)
	NewContext(ContextDisableHostCache)
	NewContext(ContextDisablePidCache)
}
