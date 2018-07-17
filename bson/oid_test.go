package bson

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestNewOid(t *testing.T) {
	fmt.Printf("gen oid: %s\n", NewOid())
}

func TestNewOidWithCtx(t *testing.T) {
	fmt.Printf("gen oid: %s\n", NewOidWithCtx(NewContext(ContextNone)))
}

func TestNewOidFromStr(t *testing.T) {
	fmt.Printf("gen oid with str: %s\n", NewOidFromStr("5b3cb2c679bc20c1037200e1"))
}

func TestNewOidFromData(t *testing.T) {
	data := make([]byte, 12)
	rand.Read(data)
	fmt.Printf("gen oid from data: %s\n", NewOidFromData(data))
}

func TestNewOidSequence(t *testing.T) {
	fmt.Printf("gen oid sequence: %s\n", NewOidSequence())
}
