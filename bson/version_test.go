package bson

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	fmt.Printf("version: %v\n", GetVersion())
	fmt.Printf("version: %d.%d.%d\n", GetMajorVersion(), GetMinorVersion(), GetMicroVersion())
}
