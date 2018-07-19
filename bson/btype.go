package bson

/*
#include <bson.h>
*/
import "C"

// BType The bson_type_t enumeration contains all of the types from the BSON Specification.
// It can be used to determine the type of a field at runtime.
// http://mongoc.org/libbson/current/bson_type_t.html
type BType C.bson_type_t

var (
	// BTypeEOD eod
	BTypeEOD = C.BSON_TYPE_EOD
	// BTypeDOUBLE double
	BTypeDOUBLE = C.BSON_TYPE_DOUBLE
	// BTypeUTF8 utf-8
	BTypeUTF8 = C.BSON_TYPE_UTF8
	// BTypeDOCUMENT document
	BTypeDOCUMENT = C.BSON_TYPE_DOCUMENT
	// BTypeARRAY array
	BTypeARRAY = C.BSON_TYPE_ARRAY
	// BTypeBINARY binary
	BTypeBINARY = C.BSON_TYPE_BINARY
	// BTypeUNDEFINED  undefined
	BTypeUNDEFINED = C.BSON_TYPE_UNDEFINED
	// BTypeOID oid
	BTypeOID = C.BSON_TYPE_OID
	// BTypeDATETIME date time
	BTypeDATETIME = C.BSON_TYPE_DATE_TIME
	// BTypeNULL null
	BTypeNULL = C.BSON_TYPE_NULL
	// BTypeREGEX regex
	BTypeREGEX = C.BSON_TYPE_REGEX
	// BTypeDBPOINTER database pointer
	BTypeDBPOINTER = C.BSON_TYPE_DBPOINTER
	// BTypeCODE code
	BTypeCODE = C.BSON_TYPE_CODE
	// BTypeSYMBOL symbol
	BTypeSYMBOL = C.BSON_TYPE_SYMBOL
	// BTypeCODEWSCOPE code ws cope
	BTypeCODEWSCOPE = C.BSON_TYPE_CODEWSCOPE
	// BTypeINT32 int32
	BTypeINT32 = C.BSON_TYPE_INT32
	// BTypeTIMESTAMP timestamp
	BTypeTIMESTAMP = C.BSON_TYPE_TIMESTAMP
	// BTypeINT64 int64
	BTypeINT64 = C.BSON_TYPE_INT64
	// BTypeDECIMAL128 decimal128
	BTypeDECIMAL128 = C.BSON_TYPE_DECIMAL128
	// BTypeMAXKEY max key
	BTypeMAXKEY = C.BSON_TYPE_MAXKEY
	// BTypeMINKEY min key
	BTypeMINKEY = C.BSON_TYPE_MINKEY
)
