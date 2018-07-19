package bson

/*
#include <bson.h>
*/
import "C"
import "unsafe"

// Free This function shall free the memory supplied by mem.
// This should be used by functions that require you free the result with bson_free().
// http://mongoc.org/libbson/current/bson_free.html
func Free(mem *interface{}) {
	C.bson_free(unsafe.Pointer(mem))
}

// FreeZero This function behaves like bson_free() except that it zeroes the memory first.
// This can be useful if you are storing passwords or other similarly important data.
// Note that if it truly is important,
// you probably want a page protected with mlock() as well to prevent it swapping to disk.
// http://mongoc.org/libbson/current/bson_zero_free.html
func FreeZero() {

}
