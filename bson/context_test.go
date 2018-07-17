package bson

import "testing"

func TestNewContext(t *testing.T) {
	NewContext(ContextFlagNone)
	NewContext(ContextFlagThreadSafe)
	NewContext(ContextFlagDisableHostCache)
	NewContext(ContextFlagDisablePidCache)
}
