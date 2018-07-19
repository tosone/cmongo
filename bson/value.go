package bson

/*
#include <bson.h>
*/
import "C"

// Value The bson_value_t structure is a boxed type for encapsulating a runtime determined type.
// http://mongoc.org/libbson/current/bson_value_t.html
type Value C.bson_value_t

// Copy This function will copy the boxed content in src into dst.
// dst must be freed with bson_value_destroy() when no longer in use.
// http://mongoc.org/libbson/current/bson_value_copy.html
func (v *Value) Copy(dst *Value) {
	C.bson_value_copy((*C.bson_value_t)(v), (*C.bson_value_t)(dst))
}

// Destroy Releases any resources associated with value. Does nothing if value is NULL.
// http://mongoc.org/libbson/current/bson_value_destroy.html
func (v *Value) Destroy() {
	C.bson_value_destroy((*C.bson_value_t)(v))
}
