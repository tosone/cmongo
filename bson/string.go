package bson

/*
#include <bson.h>
*/
import "C"

// String bson_string_t is an abstraction for building strings.
// As chunks are added to the string, allocations are performed in powers of two.
// This API is useful if you need to build UTF-8 encoded strings.
// http://mongoc.org/libbson/current/bson_string_t.html
type String C.bson_string_t

// StringNew Creates a new string builder, which uses power-of-two growth of buffers.
// Use the various bson_string_append*() functions to append to the string.
// http://mongoc.org/libbson/current/bson_string_new.html
func StringNew(str string) (s *String) {
	s = (*String)(C.bson_string_new(C.CString(str)))
	return
}

// Free Frees the bson_string_t structure and optionally the string itself.
// string->str if free_segment is false, otherwise NULL.
// http://mongoc.org/libbson/current/bson_string_free.html
func (s *String) Free(freeSegment bool) string {
	return C.GoString(C.bson_string_free((*C.bson_string_t)(s), C.bool(freeSegment)))
}

// Truncate Truncates the string so that it is len bytes in length.
// This must be smaller or equal to the current length of the string.
// A \0 byte will be placed where the end of the string occurs.
// http://mongoc.org/libbson/current/bson_string_truncate.html
func (s *String) Truncate(length uint32) {
	C.bson_string_truncate((*C.bson_string_t)(s), C.uint32_t(length))
}

// Append Appends the ASCII or UTF-8 encoded string str to string.
// This is not suitable for embedding NULLs in strings.
// http://mongoc.org/libbson/current/bson_string_append.html
func (s *String) Append(str string) {
	C.bson_string_append((*C.bson_string_t)(s), C.CString(str))
}

// AppendChar Appends c to the string builder string.
// http://mongoc.org/libbson/current/bson_string_append_c.html
func (s *String) AppendChar(char int) {
	C.bson_string_append_c((*C.bson_string_t)(s), C.char(char))
}

// AppendUniChar Appends a unicode character to string.
// The character will be encoded into its multi-byte UTF-8 representation.
// http://mongoc.org/libbson/current/bson_string_append_unichar.html
func (s *String) AppendUniChar(uniChar uint32) {
	C.bson_string_append_unichar((*C.bson_string_t)(s), C.bson_unichar_t(uniChar))
}
