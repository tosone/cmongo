package bson

/*
#include <bson.h>
*/
import "C"

// Decimal128 The bson_decimal128_t structure represents the IEEE-754 Decimal128 data type.
// http://mongoc.org/libbson/current/bson_decimal128_t.html
type Decimal128 struct {
	decimal128 *C.bson_decimal128_t
}

// String Converts dec into a printable string.
// http://mongoc.org/libbson/current/bson_decimal128_to_string.html
func (d *Decimal128) String() string {
	var c = new(C.char)
	C.bson_decimal128_to_string(d.decimal128, c)
	return C.GoString(c)
}

// ParseDecimal128 Parses the string containing ascii encoded Decimal128 and initialize the bytes in dec.
// See the Decimal128 specification for the exact string format.
// http://mongoc.org/libbson/current/bson_decimal128_from_string.html
func ParseDecimal128(num string) (d *Decimal128, b bool) {
	d = new(Decimal128)
	d.decimal128 = new(C.bson_decimal128_t)
	b = bool(C.bson_decimal128_from_string(C.CString(num), d.decimal128))
	return
}

// ParseDecimal128WLen Parses the string containing ascii encoded Decimal128 and initialize the bytes in dec.
// See the Decimal128 specification for the exact string format.
// http://mongoc.org/libbson/current/bson_decimal128_from_string_w_len.html
func ParseDecimal128WLen(num string) (d *Decimal128, b bool) {
	d = new(Decimal128)
	d.decimal128 = new(C.bson_decimal128_t)
	b = bool(C.bson_decimal128_from_string_w_len(C.CString(num), C.int(len(num)), d.decimal128))
	return
}
