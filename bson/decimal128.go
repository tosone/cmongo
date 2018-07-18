package bson

/*
#include <bson.h>
*/
import "C"

// Decimal128 The bson_decimal128_t structure represents the IEEE-754 Decimal128 data type.
// http://mongoc.org/libbson/current/bson_decimal128_t.html
type Decimal128 C.bson_decimal128_t

// String Converts dec into a printable string.
// http://mongoc.org/libbson/current/bson_decimal128_to_string.html
func (d *Decimal128) String() string {
	var c = new(C.char)
	C.bson_decimal128_to_string((*C.bson_decimal128_t)(d), c)
	return C.GoString(c)
}

// ParseDecimal128 Parses the string containing ascii encoded Decimal128 and initialize the bytes in dec.
// See the Decimal128 specification for the exact string format.
// http://mongoc.org/libbson/current/bson_decimal128_from_string.html
func ParseDecimal128(num string) (decimal *Decimal128, b bool) {
	decimal = (*Decimal128)(new(C.bson_decimal128_t))
	b = bool(C.bson_decimal128_from_string(C.CString(num), (*C.bson_decimal128_t)(decimal)))
	return
}

// ParseDecimal128WLen Parses the string containing ascii encoded Decimal128 and initialize the bytes in dec.
// See the Decimal128 specification for the exact string format.
// http://mongoc.org/libbson/current/bson_decimal128_from_string_w_len.html
func ParseDecimal128WLen(num string) (decimal *Decimal128, b bool) {
	decimal = (*Decimal128)(new(C.bson_decimal128_t))
	b = bool(C.bson_decimal128_from_string_w_len(C.CString(num), C.int(len(num)), (*C.bson_decimal128_t)(decimal)))
	return
}
