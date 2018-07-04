package bson

/*
#include <bson.h>
*/
import "C"
import "unsafe"

func NewOid() string {
	var cOid C.bson_oid_t
	var str = new(C.char)

	C.bson_oid_init(&cOid, nil)
	C.bson_oid_to_string(&cOid, str)
	return C.GoString(str)
}

func NewOidWithCtx(ctx *C.bson_context_t) string {
	var cOid C.bson_oid_t
	var str = new(C.char)

	C.bson_oid_init(&cOid, ctx)
	C.bson_oid_to_string(&cOid, str)
	return C.GoString(str)
}

func NewOidFromStr(origin string) string {
	var cOid C.bson_oid_t
	var str = new(C.char)
	var cOrigin = C.CString(origin)
	defer C.free(unsafe.Pointer(cOrigin))
	C.bson_oid_init_from_string(&cOid, cOrigin)
	C.bson_oid_to_string(&cOid, str)
	return C.GoString(str)
}

func NewOidFromData(data []byte) string {
	var cOid C.bson_oid_t
	var str = new(C.char)
	var cData = C.CBytes(data)
	defer C.free(cData)
	C.bson_oid_init_from_data(&cOid, (*C.uchar)(cData))
	C.bson_oid_to_string(&cOid, str)
	return C.GoString(str)
}

func NewOidSequence() string {
	var cOid C.bson_oid_t
	var str = new(C.char)
	C.bson_oid_init_sequence(&cOid, NewContext(ContextThreadSafe))
	C.bson_oid_to_string(&cOid, str)
	return C.GoString(str)
}
