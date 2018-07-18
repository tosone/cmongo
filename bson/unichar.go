package bson

/*
#include <bson.h>
*/
import "C"

// UniChar bson_unichar_t provides an abstraction on a single unicode character.
// It is the 32-bit representation of a character.
// As UTF-8 can contain multi-byte characters, this should be used when iterating through UTF-8 text.
// http://mongoc.org/libbson/current/bson_unichar_t.html
type UniChar C.bson_unichar_t
