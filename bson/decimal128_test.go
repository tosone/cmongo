package bson

import (
	"fmt"
	"testing"
)

func TestParseDecimal128(t *testing.T) {
	var d = new(Decimal128)
	var temp bool
	if d, temp = ParseDecimal128("100.01"); !temp {
		t.Error("cannot parse string")
	}
	fmt.Printf("decimal struct: %+v\n", d.decimal128)
	fmt.Printf("decimal128 string: %s\n", d.String())
}

func TestParseDecimal128WLen(t *testing.T) {
	var d = new(Decimal128)
	var temp bool
	if d, temp = ParseDecimal128WLen("100.01"); !temp {
		t.Error("cannot parse string")
	}
	fmt.Printf("decimal struct: %+v\n", d.decimal128)
	fmt.Printf("decimal128 string: %s\n", d.String())
}
