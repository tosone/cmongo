package bson

/*
#include <bson.h>
*/
import "C"

// Subtype This enumeration contains the various subtypes that may be used in a binary field.
// See http://bsonspec.org for more information.
// http://mongoc.org/libbson/current/bson_subtype_t.html
type Subtype C.bson_subtype_t

var (
	// SubtypeBinary binary
	SubtypeBinary = C.BSON_SUBTYPE_BINARY
	// SubtypeFunction function
	SubtypeFunction = C.BSON_SUBTYPE_FUNCTION
	// SubtypeBinaryDeprecated deprecated binary
	SubtypeBinaryDeprecated = C.BSON_SUBTYPE_BINARY_DEPRECATED
	// SubtypeUUIDDeprecated deprecated uuid
	SubtypeUUIDDeprecated = C.BSON_SUBTYPE_UUID_DEPRECATED
	// SubtypeUUID uuid
	SubtypeUUID = C.BSON_SUBTYPE_UUID
	// SubtypeMD5 md5
	SubtypeMD5 = C.BSON_SUBTYPE_MD5
	// SubtypeUser user
	SubtypeUser = C.BSON_SUBTYPE_USER
)
