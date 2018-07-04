package bson

import (
	"fmt"
	"testing"
)

func TestNewOid(t *testing.T) {
	fmt.Printf("gen oid: %s\n", NewOid())
}

func TestNewOidFromStr(t *testing.T) {
	fmt.Printf("gen oid with 'test': %s\n", NewOidFromStr("5b3cb2c679bc20c1037200e1"))
}

